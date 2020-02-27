package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetIssueComments(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	cs, err := client.GetIssueComments("12345", nil)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetIssueComments: %+v\n", cs)
}

func TestGetPullRequestComments(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	cs, err := client.GetPullRequestComments("12345", "67890", 123, nil)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetPullRequestComments: %+v\n", cs)
}
