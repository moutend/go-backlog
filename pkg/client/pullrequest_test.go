package client

import (
	"log"
	"net/url"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestAddPullRequest(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	query := url.Values{}

	query.Add("summary", "Pull request summary")
	query.Add("description", "Pull request description")
	query.Add("base", "master")
	query.Add("branch", "foo")

	pr, err := client.AddPullRequest("62794", "my-test-repo", query)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("AddPullRequest: %+v\n", pr)
}

func TestGetPullRequest(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	pr, err := client.GetPullRequest("62794", "my-test-repo", 1)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetPullRequest: %+v\n", pr)
}

func TestUpdatePullRequest(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	query := url.Values{}

	query.Add("summary", "Updated issue summary")

	pr, err := client.UpdatePullRequest("62794", "my-test-repo", 1, query)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("UpdatePullRequest: %+v\n", pr)
}

func TestGetPullRequests(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	prs, err := client.GetPullRequests("62794", "my-test-repo", nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetPullRequests: %+v\n", prs)
}

func TestGetPullRequestsCount(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	count, err := client.GetPullRequestsCount("62794", "my-test-repo", nil)

	if err != nil {
		t.Fatal(err)
	}
	log.Printf("GetPullRequestsCount: %+v\n", count)
}
