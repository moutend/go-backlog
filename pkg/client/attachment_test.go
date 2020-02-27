package client

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestPostAttachment(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	baseDir, err := ioutil.TempDir("", "testing")

	if err != nil {
		t.Fatal(err)
	}

	outputPath := filepath.Join(baseDir, "testing.txt")

	if err := ioutil.WriteFile(outputPath, []byte("Hello, World!"), 0644); err != nil {
		t.Fatal(err)
	}

	attachment, err := client.PostAttachment(outputPath)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("DONE\tPostAttachment: %+v\n", attachment)
}
