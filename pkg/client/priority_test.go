package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetPriorities(t *testing.T) {
	client, err := NewClient(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	ps, err := client.GetPriorities()

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetPriorities: %+v\n", ps)
}
