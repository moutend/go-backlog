package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetUsers(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	us, err := client.GetUsers()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetUsers: %+v\n", us)
}

func TestGetMyself(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	u, err := client.GetMyself()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetMyself: %+v\n", u)
}
