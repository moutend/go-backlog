package client

import (
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetNotifications(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	ns, err := client.GetNotifications(nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetNotifications: %+v\n", ns)
}
