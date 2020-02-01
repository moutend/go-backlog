package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

// GetIssueComments returns list of comments in issue.
//
// https://developer.nulab.com/docs/backlog/api/2/get-comment-list/#get-comment-list
func (c *Client) GetIssueComments(issueIdOrKey string, query url.Values) ([]*Comment, error) {
	return c.GetIssueCommentsContext(context.Background(), issueIdOrKey, query)
}

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

func (c *Client) GetPullRequestComments(projectKeyOrId, repositoryNameOrId, number string, query url.Values) ([]*Comment, error) {
	return c.GetPullRequestCommentsContext(context.Background(), projectKeyOrId, repositoryNameOrId, number, query)
}

func (c *Client) GetPullRequestCommentsContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, number string, query url.Values) ([]*Comment, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests", number, "comments",
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
