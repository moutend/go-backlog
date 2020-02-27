package client

import (
	"testing"

	"github.com/moutend/go-backlog/pkg/types"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestAddPullRequest(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	pullRequest := &types.PullRequest{
		ProjectId:    12345,
		RepositoryId: 67890,
		Summary:      "summary",
		Base:         "base-branch",
		Branch:       "branch",
	}

	pr, err := client.AddPullRequest(pullRequest, nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("AddPullRequest: %+v\n", pr)
}

func TestGetPullRequest(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	pr, err := client.GetPullRequest("12345", "67890", 123)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetPullRequest: %+v\n", pr)
}

func TestUpdatePullRequest(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	pullRequest := &types.PullRequest{
		ProjectId:    12345,
		RepositoryId: 67890,
		Number:       123,
		Summary:      "Updated pull request",
		Description:  "Updated description",
	}

	pr, err := client.UpdatePullRequest(pullRequest, nil, "")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("UpdatePullRequest: %+v\n", pr)
}

func TestGetPullRequestsCount(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	count, err := client.GetPullRequestsCount("12345", "67890")

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetPullRequestsCount: %+v\n", count)
}

func TestGetPullRequests(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	prs, err := client.GetPullRequests("12345", "67890", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetPullRequests: %+v\n", prs)
}

func TestGetAllPullRequests(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	prs, err := client.GetAllPullRequests("12345", "67890")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetAllPullRequests: %+v\n", prs)
}
