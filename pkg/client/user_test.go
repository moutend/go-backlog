package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetUsers(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	us, err := client.GetUsers()

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetUsers: %+v\n", us)
}

func TestGetMyself(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	u, err := client.GetMyself()

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetMyself: %+v\n", u)
}
