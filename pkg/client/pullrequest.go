package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"time"

	. "github.com/moutend/go-backlog/pkg/types"
)

// AddPullRequest Adds pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/add-pull-request/#add-pull-request
func (c *Client) AddPullRequest(pullRequest *PullRequest, notifiedUsers []*User) (*PullRequest, error) {
	return c.AddPullRequestContext(context.Background(), pullRequest, notifiedUsers)
}

// AddPullRequestContext accepts context.
func (c *Client) AddPullRequestContext(ctx context.Context, pullRequest *PullRequest, notifiedUsers []*User) (*PullRequest, error) {
	if pullRequest == nil {
		return nil, fmt.Errorf("client: pullRequest is required")
	}

	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, fmt.Sprint(pullRequest.ProjectId),
		"git", "repositories", fmt.Sprint(pullRequest.RepositoryId),
		"pullRequests",
	))

	if err != nil {
		return nil, err
	}

	query := pullRequest.EncodeQuery()

	if len(notifiedUsers) > 0 {
		for _, notifiedUser := range notifiedUsers {
			query.Add("notifiedUserId", fmt.Sprint(notifiedUser.Id))
		}
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

	var createdPullRequest *PullRequest

	if err := json.Unmarshal(body, &createdPullRequest); err != nil {
		return nil, err
	}

	return createdPullRequest, nil
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
func (c *Client) UpdatePullRequest(pullRequest *PullRequest, notifiedUsers []*User, comment string) (*PullRequest, error) {
	return c.UpdatePullRequestContext(context.Background(), pullRequest, notifiedUsers, comment)
}

// UpdatePullRequestContext accepts context.
func (c *Client) UpdatePullRequestContext(ctx context.Context, pullRequest *PullRequest, notifiedUsers []*User, comment string) (*PullRequest, error) {
	if pullRequest == nil {
		return nil, fmt.Errorf("client: pullRequest is required")
	}

	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, fmt.Sprint(pullRequest.ProjectId),
		"git", "repositories", fmt.Sprint(pullRequest.RepositoryId),
		"pullRequests", fmt.Sprint(pullRequest.Number),
	))

	if err != nil {
		return nil, err
	}

	query := url.Values{}

	if pullRequest.Summary != "" {
		query.Add("summary", pullRequest.Summary)
	}
	if pullRequest.Description != "" {
		query.Add("description", pullRequest.Description)
	}
	if pullRequest.Issue != nil {
		query.Add("description", fmt.Sprint(pullRequest.Issue.Id))
	}
	if pullRequest.Assignee != nil {
		query.Add("assigneeId", fmt.Sprint(pullRequest.Assignee.Id))
	}
	if len(notifiedUsers) > 0 {
		for _, notifiedUser := range notifiedUsers {
			query.Add("notifiedUserId", fmt.Sprint(notifiedUser.Id))
		}
	}
	if comment != "" {
		query.Add("comment", comment)
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

	var updatedPullRequest *PullRequest

	if err := json.Unmarshal(body, &updatedPullRequest); err != nil {
		return nil, err
	}

	return updatedPullRequest, nil
}

// GetPullRequestsCount returns number of pull requests.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-number-of-pull-requests/#get-number-of-pull-requests
func (c *Client) GetPullRequestsCount(projectIdOrKey, repositoryIdOrName string) (int64, error) {
	return c.GetPullRequestsCountContext(context.Background(), projectIdOrKey, repositoryIdOrName)
}

// GetPullRequestsCountContext accepts context.
func (c *Client) GetPullRequestsCountContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string) (int64, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectIdOrKey,
		"git", "repositories", repositoryIdOrName,
		"pullRequests", "count",
	))

	if err != nil {
		return -1, err
	}

	res, err := c.getContext(ctx, path, nil)

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

// GetAllPullRequests returns all pull requests.
func (c *Client) GetAllPullRequests(projectIdOrKey, repositoryIdOrName string) ([]*PullRequest, error) {
	return c.GetAllPullRequestsContext(context.Background(), projectIdOrKey, repositoryIdOrName)
}

// GetAllPullRequestsContext accepts context.
func (c *Client) GetAllPullRequestsContext(ctx context.Context, projectIdOrKey, repositoryIdOrName string) ([]*PullRequest, error) {
	count, err := c.GetPullRequestsCount(projectIdOrKey, repositoryIdOrName)

	if err != nil {
	}
	allPullRequests := []*PullRequest{}

	offset := 0
	times := int((count / 100) + 1)

	for i := 0; i < times; i++ {
		query := url.Values{}

		query.Add("count", "100")
		query.Add("offset", fmt.Sprint(offset))

		pullRequests, err := c.GetPullRequestsContext(ctx, projectIdOrKey, repositoryIdOrName, query)

		if err != nil {
			return allPullRequests, err
		}

		allPullRequests = append(allPullRequests, pullRequests...)

		time.Sleep(250 * time.Millisecond)

		offset += 100
	}

	return allPullRequests, nil
}
