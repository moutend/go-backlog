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

// GetIssueComments returns list of comments in issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-comment-list/#get-comment-list
func (c *Client) GetIssueComments(issueIdOrKey string, query url.Values) ([]*Comment, error) {
	return c.GetIssueCommentsContext(context.Background(), issueIdOrKey, query)
}

// GetIssueCommentsContext accepts context.
func (c *Client) GetIssueCommentsContext(ctx context.Context, issueIdOrKey string, query url.Values) ([]*Comment, error) {
	path, err := c.root.Parse(path.Join(V2IssuesPath, issueIdOrKey, "comments"))

	if err != nil {
		return nil, err
	}

	res, err := c.getContext(ctx, path, query)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var comments []*Comment

	if err := json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetPullRequestComments returns list of pull request comments.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-pull-request-comment/
func (c *Client) GetPullRequestComments(projectKeyOrId, repositoryNameOrId string, number int64, query url.Values) ([]*Comment, error) {
	return c.GetPullRequestCommentsContext(context.Background(), projectKeyOrId, repositoryNameOrId, number, query)
}

// GetPullRequestCommentsContext accepts context.
func (c *Client) GetPullRequestCommentsContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, number int64, query url.Values) ([]*Comment, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests", fmt.Sprint(number), "comments",
	))
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

	var comments []*Comment

	if err := json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}
