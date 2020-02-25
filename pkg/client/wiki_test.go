package client

import (
	"log"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
	. "github.com/moutend/go-backlog/pkg/types"
)

func TestGetWikis(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	ws, err := client.GetWikis("62794", nil)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWikis: %+v\n", ws)
}

func TestAddWiki(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	wiki := &Wiki{
		Name:      "wiki name",
		Content:   "wiki content",
		ProjectId: 1234567,
	}

	createdWiki, err := client.AddWiki(wiki, false)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("AddWiki: %+v\n", createdWiki)
}

func TestGetWiki(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	w, err := client.GetWiki(312145)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWiki: %+v\n", w)
}

func TestUpdateWiki(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	wiki := &Wiki{
		Id:        7777777,
		Name:      "Wiki name",
		Content:   "Wiki content",
		ProjectId: 1234567,
	}

	updatedWiki, err := client.UpdateWiki(wiki, false)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("UpdateWiki: %+v\n", updatedWiki)
}

func TestDeleteWiki(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	deletedWiki, err := client.DeleteWiki(7777777)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("DeleteWiki: %+v\n", deletedWiki)
}

func TestGetWikiCount(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`{}`), testutil.EnableHTTPRequest))

	count, err := client.GetWikiCount("projectKey")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWikiCount: %+v\n", count)
}

func TestGetWikiTags(t *testing.T) {
	client, err := New(testutil.BacklogSpace, testutil.BacklogToken)

	if err != nil {
		t.Fatal(err)
	}

	client.SetHTTPClient(testutil.NewTestClient([]byte(`[]`), testutil.EnableHTTPRequest))

	tags, err := client.GetWikiTags("projectKey")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("GetWikiTags: %+v\n", tags)
}
