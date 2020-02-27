package client

import (
	"testing"

	. "github.com/moutend/go-backlog/pkg/types"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestGetWikis(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	wikis, err := client.GetWikis("projectKey", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetWikis: %+v\n", wikis)
}

func TestAddWiki(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	wiki := &Wiki{
		Name:      "wiki name",
		Content:   "wiki content",
		ProjectId: 1234567,
	}

	createdWiki, err := client.AddWiki(wiki, false)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("AddWiki: %+v\n", createdWiki)
}

func TestGetWiki(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	wiki, err := client.GetWiki(12345)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetWiki: %+v\n", wiki)
}

func TestUpdateWiki(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	wiki := &Wiki{
		Id:        12345,
		Name:      "Wiki name",
		Content:   "Wiki content",
		ProjectId: 1234567,
	}

	updatedWiki, err := client.UpdateWiki(wiki, false)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("UpdateWiki: %+v\n", updatedWiki)
}

func TestDeleteWiki(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	deletedWiki, err := client.DeleteWiki(12345)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("DeleteWiki: %+v\n", deletedWiki)
}

func TestGetWikiCount(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	count, err := client.GetWikiCount("projectKey")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetWikiCount: %+v\n", count)
}

func TestGetWikiTags(t *testing.T) {
	client, err := New("test.backlog.com", "token", OptionHTTPClient(testutil.NewFakeClient(t)))

	if err != nil {
		t.Fatal(err)
	}

	tags, err := client.GetWikiTags("projectKey")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetWikiTags: %+v\n", tags)
}
