package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetNotifications(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	ns, err := client.GetNotifications(nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetNotifications: %+v\n", ns)
}
