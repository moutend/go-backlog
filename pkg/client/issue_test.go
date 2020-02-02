package client

import (
	"log"
	"net/url"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
	. "github.com/moutend/go-backlog/pkg/types"
)

func TestAddIssue(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	projectId := uint64(62794)

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

	log.Printf("AddIssue: %+v\n", createdIssue)
}

func TestGetIssue(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	i, err := client.GetIssue("3395955")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetIssue: %+v\n", i)
}

func TestUpdateIssue(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	query := url.Values{}

	query.Add("summary", "Updated issue summary")

	i, err := client.UpdateIssue("3395955", query)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("UpdateIssue: %+v\n", i)
}

func TestDeleteIssue(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	i, err := client.DeleteIssue("3395535")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("DeleteIssue: %+v\n", i)
}

func TestGetIssues(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	query := url.Values{}

	query.Add("count", "100")
	query.Add("projectId[]", "62794")

	is, err := client.GetIssues(query)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetIssues: %+v\n", is)
}

func TestGetIssuesCount(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	count, err := client.GetIssuesCount(nil)

	if err != nil {
		t.Fatal(err)
	}
	log.Printf("GetIssuesCount: %+v\n", count)
}

func TestGetIssueTypes(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	its, err := client.GetIssueTypes("62794")

	if err != nil {
		t.Fatal(err)
	}
	log.Printf("GetIssueTypes: %+v\n", its)
}
