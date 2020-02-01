package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2IssuesPath      = "/api/v2/issues"
	V2IssuesCountPath = "/api/v2/issues/count"
)

// AddIssue adds new issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/add-issue/#add-issue
func (c *Client) AddIssue(query url.Values) (*Issue, error) {
	return c.AddIssueContext(context.Background(), query)
}

// AddIssueContext accepts context.
func (c *Client) AddIssueContext(ctx context.Context, query url.Values) (*Issue, error) {
	path, err := c.root.Parse(V2IssuesPath)

	if err != nil {
		return nil, err
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

	var issue *Issue

	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}

	return issue, nil
}

// GetIssue returns information about issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-issue/#get-issue
func (c *Client) GetIssue(issueKeyOrId string) (*Issue, error) {
	return c.GetIssueContext(context.Background(), issueKeyOrId)
}

// GetIssueContext accepts context.
func (c *Client) GetIssueContext(ctx context.Context, issueKeyOrId string) (*Issue, error) {
	path, err := c.root.Parse(path.Join(V2IssuesPath, issueKeyOrId))

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

	var issue *Issue

	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}

	return issue, nil
}

// UpdateIssue updates information about issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/update-issue/#update-issue
func (c *Client) UpdateIssue(issueKeyOrId string, query url.Values) (*Issue, error) {
	return c.UpdateIssueContext(context.Background(), issueKeyOrId, query)
}

// UpdateIssueContext accepts context.
func (c *Client) UpdateIssueContext(ctx context.Context, issueKeyOrId string, query url.Values) (*Issue, error) {
	path, err := c.root.Parse(path.Join(V2IssuesPath, issueKeyOrId))

	if err != nil {
		return nil, err
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

	var issue *Issue

	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}

	return issue, nil
}

// DeleteIssue deletes issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/delete-issue/#delete-issue
func (c *Client) DeleteIssue(issueKeyOrId string) (*Issue, error) {
	return c.DeleteIssueContext(context.Background(), issueKeyOrId)
}

// DeleteIssueContext accepts context.
func (c *Client) DeleteIssueContext(ctx context.Context, issueKeyOrId string) (*Issue, error) {
	path, err := c.root.Parse(path.Join(V2IssuesPath, issueKeyOrId))

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

	var issue *Issue

	if err := json.Unmarshal(body, &issue); err != nil {
		return nil, err
	}

	return issue, nil
}

// GetIssues returns list of issues.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-issue-list/#get-issue-list
func (c *Client) GetIssues(query url.Values) ([]*Issue, error) {
	return c.GetIssuesContext(context.Background(), query)
}

// GetIssuesContext accepts context.
func (c *Client) GetIssuesContext(ctx context.Context, query url.Values) ([]*Issue, error) {
	path, err := c.root.Parse(V2IssuesPath)

	if err != nil {
		return nil, err
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

	var issues []*Issue

	if err = json.Unmarshal(body, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// GetIssuesCount returns number of issues.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/count-issue/#count-issue
func (c *Client) GetIssuesCount(query url.Values) (int64, error) {
	return c.GetIssuesCountContext(context.Background(), query)
}

// GetIssuesCountContext accepts context.
func (c *Client) GetIssuesCountContext(ctx context.Context, query url.Values) (int64, error) {
	path, err := c.root.Parse(V2IssuesCountPath)

	if err != nil {
		return -1, err
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

	var v struct {
		Count int64 `json:"count"`
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return -1, err
	}

	return v.Count, nil
}

// GetIssueTypes returns list of Issue Types in the project.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-issue-type-list/#get-issue-type-list
func (c *Client) GetIssueTypes(projectIdOrKey string) ([]*IssueType, error) {
	return c.GetIssueTypesContext(context.Background(), projectIdOrKey)
}

// GetIssueTypesContext accepts context.
func (c *Client) GetIssueTypesContext(ctx context.Context, projectIdOrKey string) ([]*IssueType, error) {
	path, err := c.root.Parse(path.Join(V2ProjectsPath, projectIdOrKey, "issueTypes"))

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

	var issueTypes []*IssueType

	if err := json.Unmarshal(body, &issueTypes); err != nil {
		return nil, err
	}

	return issueTypes, nil
}
