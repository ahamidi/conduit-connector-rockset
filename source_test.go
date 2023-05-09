package rockset_test

import (
	"context"
	"testing"

	rockset "github.com/ahmeroxa/conduit-connector-rockset"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	con := rockset.NewSource()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
