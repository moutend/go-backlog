package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetPriorities(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	ps, err := client.GetPriorities()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetPriorities: %+v\n", ps)
}
