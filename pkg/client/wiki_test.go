package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetWikis(t *testing.T) {
	client, err := NewClient(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	ws, err := client.GetWikis("62794", nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWikis: %+v\n", ws)
}

func TestGetWiki(t *testing.T) {
	client, err := NewClient(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	w, err := client.GetWiki(312145)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWiki: %+v\n", w)
}
