package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	rockset "github.com/ahmeroxa/conduit-connector-rockset"
)

func main() {
	sdk.Serve(rockset.Connector)
}
