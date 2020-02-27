package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetProjectStatuses(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	pss, err := client.GetProjectStatuses("12345")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetProjectStatuses: %+v\n", pss)
}
func TestGetProject(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	p, err := client.GetProject("12345")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetProject: %+v\n", p)
}

func TestGetProjects(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	ps, err := client.GetProjects(nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetProjects: %+v\n", ps)
}
