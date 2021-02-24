package types

import (
	"fmt"
	"net/url"
)

// PullRequest represents the pull request.
type PullRequest struct {
	Id           int64   `json:"id"`
	ProjectId    int64   `json:"projectId"`
	RepositoryId int64   `json:"repositoryID"`
	Number       int     `json:"number"`
	Summary      string  `json:"summary"`
	Description  string  `json:"description"`
	Base         string  `json:"base"`
	Branch       string  `json:"branch"`
	Status       *Status `json:"status"`
	Assignee     *User   `json:"assignee"`
	Issue        *Issue  `json:"issue"`
	BaseCommit   string  `json:"baseCommit"`
	BranchCommit string  `json:"branchCommit"`
	CloseAt      *Date   `json:"closeAt"`
	MergeAt      *Date   `json:"mergeAt"`
	CreatedUser  *User   `json:"createdUser"`
	Created      *Date   `json:"created"`
	UpdatedUser  *User   `json:"updatedUser"`
	Updated      *Date   `json:"update"`
}

func (p *PullRequest) EncodeQuery() url.Values {
	query := url.Values{}

	query.Add("summary", p.Summary)
	query.Add("description", p.Description)
	query.Add("base", p.Base)
	query.Add("branch", p.Branch)

	if p.Issue != nil {
		query.Add("issueId", fmt.Sprint(p.Issue.Id))
	}
	if p.Assignee != nil {
		query.Add("assigneeId", fmt.Sprint(p.Assignee.Id))
	}

	return query
}
