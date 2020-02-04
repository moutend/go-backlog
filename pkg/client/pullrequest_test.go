package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/pkg/types"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestAddPullRequest(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	pullRequest := &types.PullRequest{
		ProjectId:    12345,
		RepositoryId: 12345,
		Summary:      "summary",
		Base:         "base-branch",
		Branch:       "branch",
	}

	pr, err := client.AddPullRequest(pullRequest, nil)

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

	pullRequest := &types.PullRequest{
		Summary:     "Updated pull request",
		Description: "Updated description",
	}

	pr, err := client.UpdatePullRequest(pullRequest, nil, "")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("UpdatePullRequest: %+v\n", pr)
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

func TestGetAllPullRequests(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	prs, err := client.GetAllPullRequests("PRJ", "my-test-repo")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetAllPullRequests: %+v\n", prs)
}
