package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetProjectStatuses(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	pss, err := client.GetProjectStatuses("62794")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetProjectStatuses: %+v\n", pss)
}
func TestGetProject(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	p, err := client.GetProject("62794")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetProject: %+v\n", p)
}

func TestGetProjects(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	ps, err := client.GetProjects(nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetProjects: %+v\n", ps)
}
