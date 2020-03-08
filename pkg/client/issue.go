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

// AddIssue adds new issue.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/add-issue/#add-issue
func (c *Client) AddIssue(issue *Issue, notifiedUsers []*User) (*Issue, error) {
	return c.AddIssueContext(context.Background(), issue, notifiedUsers)
}

// AddIssueContext accepts context.
func (c *Client) AddIssueContext(ctx context.Context, issue *Issue, notifiedUsers []*User) (*Issue, error) {
	path, err := c.root.Parse(V2IssuesPath)

	if err != nil {
		return nil, err
	}
	if issue == nil {
		return nil, fmt.Errorf("issue is required")
	}
	if issue.ProjectId == nil {
		return nil, fmt.Errorf("issue.ProjectId is required")
	}
	if issue.Summary == "" {
		return nil, fmt.Errorf("issue.Summary is required")
	}
	if issue.IssueType == nil {
		return nil, fmt.Errorf("issue.IssueType is required")
	}
	if issue.Priority == nil {
		return nil, fmt.Errorf("issue.Priority is required")
	}

	query := issue.URLValues()
	query.Del("statusId")
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

	var createdIssue *Issue

	if err := json.Unmarshal(body, &createdIssue); err != nil {
		return nil, err
	}

	return createdIssue, nil
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
func (c *Client) UpdateIssue(issue *Issue, notifiedUsers []*User, comment string) (*Issue, error) {
	return c.UpdateIssueContext(context.Background(), issue, notifiedUsers, comment)
}

// UpdateIssueContext accepts context.
func (c *Client) UpdateIssueContext(ctx context.Context, issue *Issue, notifiedUsers []*User, comment string) (*Issue, error) {
	if issue == nil {
		return nil, fmt.Errorf("issue is required")
	}

	issueIdOrKey := fmt.Sprint(issue.Id)

	if issue.IssueKey != "" {
		issueIdOrKey = issue.IssueKey
	}

	path, err := c.root.Parse(path.Join(V2IssuesPath, issueIdOrKey))

	if err != nil {
		return nil, err
	}

	query := issue.URLValues()
	query.Del("projectId")

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

	var updatedIssue *Issue

	if err := json.Unmarshal(body, &updatedIssue); err != nil {
		return nil, err
	}

	return updatedIssue, nil
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

// GetAllIssues returns all issues.
func (c *Client) GetAllIssues(maxIssues int, query url.Values) ([]*Issue, error) {
	return c.GetAllIssuesContext(context.Background(), maxIssues, query)
}

// GetAllIssuesContext accepts context.
func (c *Client) GetAllIssuesContext(ctx context.Context, maxIssues int, query url.Values) ([]*Issue, error) {
	foundIssues, err := c.GetIssuesCount(nil)

	if err != nil {
		return nil, err
	}

	allIssues := []*Issue{}

	if foundIssues == 0 {
		return allIssues, nil
	}

	issuesCount := 0
	times := int((foundIssues / 100) + 1)

	for i := 0; i < times; i++ {
		if maxIssues > 0 && issuesCount >= maxIssues {
			break
		}

		q := url.Values{}

		q.Add("count", "100")
		q.Add("offset", fmt.Sprint(issuesCount))

		for k, vs := range query {
			for _, v := range vs {
				q.Add(k, v)
			}
		}

		issues, err := c.GetIssuesContext(ctx, q)

		if err != nil {
			return allIssues, err
		}

		issuesCount += len(issues)
		allIssues = append(allIssues, issues...)

		// backlog API doesn't provide rate / limit information, sleep 250 ms to prevent sending many requests at once.
		time.Sleep(250 * time.Millisecond)
	}

	return allIssues, nil
}
