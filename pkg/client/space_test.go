package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetSpace(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	s, err := client.GetSpace()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetSpace: %+v\n", s)
}
