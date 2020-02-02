package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetIssueComments(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	cs, err := client.GetIssueComments("3395955", nil)

	if err != nil {
		t.Fatal(err)
	}
	log.Printf("GetIssueComments: %+v\n", cs)
}

func TestGetPullRequestComments(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	cs, err := client.GetPullRequestComments("62794", "my-test-repo", "1", nil)

	if err != nil {
		t.Fatal(err)
	}
	log.Printf("GetPullRequestComments: %+v\n", cs)
}
