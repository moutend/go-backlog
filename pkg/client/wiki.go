package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2WikisPath = "/api/v2/wikis"
)

// GetWiki returns information about Wiki page.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-wiki-page/
func (c *Client) GetWiki(wikiId uint64) (*Wiki, error) {
	return c.GetWikiContext(context.Background(), wikiId)
}

// GetWikiContext accepts context.
func (c *Client) GetWikiContext(ctx context.Context, wikiId uint64) (*Wiki, error) {
	path, err := c.root.Parse(path.Join(V2WikisPath, fmt.Sprint(wikiId)))

	if err != nil {
		return nil, err
	}

	res, err := c.getContext(ctx, path, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var wiki *Wiki

	if err := json.Unmarshal(body, &wiki); err != nil {
		return nil, err
	}

	return wiki, nil
}

// GetWikis returns list of Wiki pages.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-wiki-page-list/#get-wiki-page-list
func (c *Client) GetWikis(projectIdOrKey string, query url.Values) ([]*Wiki, error) {
	return c.GetWikisContext(context.Background(), projectIdOrKey, query)
}

// GetWikisContext accepts context.
func (c *Client) GetWikisContext(ctx context.Context, projectIdOrKey string, query url.Values) ([]*Wiki, error) {
	path, err := c.root.Parse(V2WikisPath)

	if err != nil {
		return nil, err
	}

	if query == nil {
		query = url.Values{}
	}
	if query.Get("projectIdOrKey") == "" {
		query.Add("projectIdOrKey", projectIdOrKey)
	}

	res, err := c.getContext(ctx, path, query)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var wikis []*Wiki

	if err := json.Unmarshal(body, &wikis); err != nil {
		return nil, err
	}

	return wikis, nil
}
