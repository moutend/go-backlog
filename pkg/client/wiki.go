package client

import (
	"bytes"
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

// AddWiki adds new Wiki page.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/add-wiki-page/#add-wiki-page
func (c *Client) AddWiki(wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return c.AddWikiContext(context.Background(), wiki, mailNotify)
}

// AddWikiContext accepts context.
func (c *Client) AddWikiContext(ctx context.Context, wiki *Wiki, mailNotify bool) (*Wiki, error) {
	path, err := c.root.Parse(V2WikisPath)

	if err != nil {
		return nil, err
	}
	if wiki == nil {
		return nil, fmt.Errorf("wiki is required")
	}

	query := url.Values{}

	query.Add("projectId", fmt.Sprint(wiki.ProjectId))
	query.Add("name", fmt.Sprint(wiki.Name))
	query.Add("content", fmt.Sprint(wiki.Content))

	if mailNotify {
		query.Add("mailNotify", "true")
	}

	payload := bytes.NewBufferString(query.Encode())

	res, err := c.postContext(ctx, path, nil, payload)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var createdWiki *Wiki

	if err := json.Unmarshal(body, &createdWiki); err != nil {
		return nil, err
	}

	return createdWiki, nil
}

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

// UpdateWiki updates information about Wiki page.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/update-wiki-page/#update-wiki-page
func (c *Client) UpdateWiki(wiki *Wiki, mailNotify bool) (*Wiki, error) {
	return c.UpdateWikiContext(context.Background(), wiki, mailNotify)
}

// UpdateWikiContext accepts context.
func (c *Client) UpdateWikiContext(ctx context.Context, wiki *Wiki, mailNotify bool) (*Wiki, error) {
	path, err := c.root.Parse(path.Join(V2WikisPath, fmt.Sprint(wiki.Id)))

	if err != nil {
		return nil, err
	}
	if wiki == nil {
		return nil, fmt.Errorf("wiki is required")
	}

	query := url.Values{}

	query.Add("name", fmt.Sprint(wiki.Name))
	query.Add("content", fmt.Sprint(wiki.Content))

	if mailNotify {
		query.Add("mailNotify", "true")
	}

	payload := bytes.NewBufferString(query.Encode())

	res, err := c.patchContext(ctx, path, nil, payload)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var createdWiki *Wiki

	if err := json.Unmarshal(body, &createdWiki); err != nil {
		return nil, err
	}

	return createdWiki, nil
}

// DeleteWiki deletes wiki.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/delete-wiki-page/#delete-wiki-page
func (c *Client) DeleteWiki(wikiId uint64) (*Wiki, error) {
	return c.DeleteWikiContext(context.Background(), wikiId)
}

// DeleteWikiContext accepts context.
func (c *Client) DeleteWikiContext(ctx context.Context, wikiId uint64) (*Wiki, error) {
	path, err := c.root.Parse(path.Join(V2WikisPath, fmt.Sprint(wikiId)))

	if err != nil {
		return nil, err
	}

	res, err := c.deleteContext(ctx, path, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var deletedWiki *Wiki

	if err := json.Unmarshal(body, &deletedWiki); err != nil {
		return nil, err
	}

	return deletedWiki, nil
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

// GetWikiCount returns number of Wiki pages.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/count-wiki-page/#count-wiki-page
func (c *Client) GetWikiCount(projectIdOrKey string) (int, error) {
	return c.GetWikiCountContext(context.Background(), projectIdOrKey)
}

// GetWikiCountContext accepts context.
func (c *Client) GetWikiCountContext(ctx context.Context, projectIdOrKey string) (int, error) {
	path, err := c.root.Parse(path.Join(V2WikisPath, "count"))

	if err != nil {
		return -1, err
	}

	query := url.Values{}

	if projectIdOrKey != "" {
		query.Add("projectIdOrKey", projectIdOrKey)
	}

	res, err := c.getContext(ctx, path, query)

	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return -1, err
	}

	var result struct {
		Count int `json:"count"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return -1, err
	}

	return result.Count, nil
}

// GetWikiTags returns list of tags that are used in the project.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-wiki-page-tag-list/#get-wiki-page-tag-list
func (c *Client) GetWikiTags(projectIdOrKey string) ([]*Tag, error) {
	return c.GetWikiTagsContext(context.Background(), projectIdOrKey)
}

// GetWikiTagsContext accepts context.
func (c *Client) GetWikiTagsContext(ctx context.Context, projectIdOrKey string) ([]*Tag, error) {
	path, err := c.root.Parse(path.Join(V2WikisPath, "tags"))

	if err != nil {
		return nil, err
	}

	query := url.Values{}

	if projectIdOrKey != "" {
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

	var tags []*Tag

	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
