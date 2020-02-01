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

// AddPullRequest Adds pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/add-pull-request/#add-pull-request
func (c *Client) AddPullRequest(projectIdOrKey, repositoryIdOrKey string, query url.Values) (*PullRequest, error) {
	return c.AddPullRequestContext(context.Background(), projectIdOrKey, repositoryIdOrKey, query)
}

// AddPullRequestContext accepts context.
func (c *Client) AddPullRequestContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, query url.Values) (*PullRequest, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests",
	))

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

	var pullRequest *PullRequest

	if err := json.Unmarshal(body, &pullRequest); err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// GetPullRequest returns pull reuqest.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-pull-request/#get-pull-request
func (c *Client) GetPullRequest(projectIdOrKey, repositoryIdOrName string, number int64) (*PullRequest, error) {
	return c.GetPullRequestContext(context.Background(), projectIdOrKey, repositoryIdOrName, number)
}

// GetPullRequestContext accepts context.
func (c *Client) GetPullRequestContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, number int64) (*PullRequest, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests", fmt.Sprint(number),
	))

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

	var pr *PullRequest

	if err := json.Unmarshal(body, &pr); err != nil {
		return nil, err
	}

	return pr, nil
}

// UpdatePullRequest updates pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/update-pull-request/#update-pull-request
func (c *Client) UpdatePullRequest(projectIdOrKey, repositoryIdOrName string, number int64, query url.Values) (*PullRequest, error) {
	return c.UpdatePullRequestContext(context.Background(), projectIdOrKey, repositoryIdOrName, number, query)
}

// UpdatePullRequestContext accepts context.
func (c *Client) UpdatePullRequestContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, number int64, query url.Values) (*PullRequest, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests", fmt.Sprint(number),
	))

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

	var pullRequest *PullRequest

	if err := json.Unmarshal(body, &pullRequest); err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// GetPullRequests returns list of pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-pull-request-list/#get-pull-request-list
func (c *Client) GetPullRequests(projectIdOrKey, repositoryIdOrName string, query url.Values) ([]*PullRequest, error) {
	return c.GetPullRequestsContext(context.Background(), projectIdOrKey, repositoryIdOrName, query)
}

// GetPullRequestsContext accepts context.
func (c *Client) GetPullRequestsContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, query url.Values) ([]*PullRequest, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests"))

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

	var prs []*PullRequest

	if err := json.Unmarshal(body, &prs); err != nil {
		return nil, err
	}

	return prs, nil
}

// GetPullRequestsCount returns number of pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-number-of-pull-requests/#get-number-of-pull-requests
func (c *Client) GetPullRequestsCount(projectIdOrKey, repositoryIdOrName string, query url.Values) (int64, error) {
	return c.GetPullRequestsCountContext(context.Background(), projectIdOrKey, repositoryIdOrName, query)
}

// GetPullRequestsCountContext accepts context.
func (c *Client) GetPullRequestsCountContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string, query url.Values) (int64, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests", "count",
	))

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
