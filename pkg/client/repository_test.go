package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetRepositories(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	rs, err := client.GetRepositories("12345", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetRepositories: %+v\n", rs)
}

func TestGetRepository(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	r, err := client.GetRepository("12345", "67890")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetRepository: %+v\n", r)
}
