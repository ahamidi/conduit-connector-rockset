package rockset

//go:generate paramgen -output=paramgen_dest.go DestinationConfig

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/rockset/rockset-go-client"
)

type Destination struct {
	sdk.UnimplementedDestination

	client     *rockset.RockClient
	workspace  string
	collection string
	config     DestinationConfig
}

type DestinationConfig struct {
	// Config includes parameters that are the same in the source and destination.
	Config
}

func NewDestination() sdk.Destination {
	// Create Destination and wrap it in the default middleware.
	return sdk.DestinationWithMiddleware(&Destination{})
}

func (d *Destination) Parameters() map[string]sdk.Parameter {
	// Parameters is a map of named Parameters that describe how to configure
	// the Destination. Parameters can be generated from DestinationConfig with
	// paramgen.
	return d.config.Parameters()
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	// Configure is the first function to be called in a connector. It provides
	// the connector with the configuration that can be validated and stored.
	// In case the configuration is not valid it should return an error.
	// Testing if your connector can reach the configured data source should be
	// done in Open, not in Configure.
	// The SDK will validate the configuration and populate default values
	// before calling Configure. If you need to do more complex validations you
	// can do them manually here.

	// map region to server url
	cfg["region"] = regionToURLMap[cfg["region"]]

	sdk.Logger(ctx).Info().Msg("Configuring Destination...")
	err := sdk.Util.ParseConfig(cfg, &d.config)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	// Open is called after Configure to signal the plugin it can prepare to
	// start writing records. If needed, the plugin should open connections in
	// this function.
	client, err := rockset.NewClient(rockset.WithAPIKey(d.config.APIKey), rockset.WithAPIServer(d.config.Region))
	if err != nil {
		log.Printf("failed to create rockset client: %v", err)
		return err
	}
	d.client = client
	return nil
}

func (d *Destination) Write(ctx context.Context, records []sdk.Record) (int, error) {
	// Write writes len(r) records from r to the destination right away without
	// caching. It should return the number of records written from r
	// (0 <= n <= len(r)) and any error encountered that caused the write to
	// stop early. Write must return a non-nil error if it returns n < len(r).

	docs := make([]interface{}, len(records))
	for i, rec := range records {
		var doc map[string]interface{}
		err := json.Unmarshal(rec.Bytes(), &doc)
		if err != nil {
			// do nothing
		}
		docs[i] = doc
	}

	//todo: handle response error codes/rate limits: https://rockset.com/docs/write-api/#write-api-limits
	statuses, err := d.client.AddDocuments(ctx, d.config.Workspace, d.config.Collection, docs)
	if err != nil {
		log.Printf("failed to write documents: %v", err)
		return 0, fmt.Errorf("failed to write documents: %v", err)
	}

	for _, status := range statuses {
		if status.HasError() {
			log.Printf("status error: %v", status.Error)
			return 0, fmt.Errorf("failed to write document(%s): %v", *status.Id, status.Error)
		}
	}

	return len(records), nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	// Teardown signals to the plugin that all records were written and there
	// will be no more calls to any other function. After Teardown returns, the
	// plugin should be ready for a graceful shutdown.
	return nil
}
