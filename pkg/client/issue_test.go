package client

import (
	"net/url"
	"testing"
	"time"

	"github.com/moutend/go-backlog/internal/testutil"
	. "github.com/moutend/go-backlog/pkg/types"
)

func TestAddIssue(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	projectId := uint64(12345)

	issue := &Issue{
		Summary:   "Issue summary",
		ProjectId: &projectId,
		IssueType: &IssueType{
			Id: 290177,
		},
		Priority: &Priority{
			Id: 4,
		},
	}

	createdIssue, err := client.AddIssue(issue, nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("AddIssue: %+v\n", createdIssue)
}

func TestGetIssue(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	i, err := client.GetIssue("12345")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetIssue: %+v\n", i)
}

func TestUpdateIssue(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	projectId := uint64(12345)

	issue := &Issue{
		Id:        12345,
		Summary:   "Updated issue summary " + time.Now().Format("15:04:05"),
		ProjectId: &projectId,
		IssueType: &IssueType{
			Id: 290177,
		},
	}

	i, err := client.UpdateIssue(issue, nil, "")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("UpdateIssue: %+v\n", i)
}

func TestDeleteIssue(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	i, err := client.DeleteIssue("12345")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("DeleteIssue: %+v\n", i)
}

func TestGetIssueTypes(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	its, err := client.GetIssueTypes("12345")

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetIssueTypes: %+v\n", its)
}

func TestGetIssuesCount(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	count, err := client.GetIssuesCount(nil)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetIssuesCount: %+v\n", count)
}

func TestGetIssues(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	query := url.Values{}

	query.Add("count", "100")
	query.Add("projectId[]", "12345")

	is, err := client.GetIssues(query)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetIssues: %+v\n", is)
}

func TestGetAllIssues(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	is, err := client.GetAllIssues(123, nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetAllIssues: %+v\n", is)
}
