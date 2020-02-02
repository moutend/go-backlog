package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetSpace(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	s, err := client.GetSpace()

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetSpace: %+v\n", s)
}
