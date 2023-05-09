package rockset_test

import (
	"context"
	"testing"

	rockset "github.com/ahmeroxa/conduit-connector-rockset"
)

func TestTeardown_NoOpen(t *testing.T) {
	con := rockset.NewDestination()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
