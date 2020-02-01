package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetRepositories(t *testing.T) {
	client, err := NewClient(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	rs, err := client.GetRepositories("LIFE", nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetRepositories: %+v\n", rs)
}

func TestGetRepository(t *testing.T) {
	client, err := NewClient(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	r, err := client.GetRepository("LIFE", "my-test-repo", nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetRepository: %+v\n", r)
}
