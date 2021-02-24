package types

import (
	"fmt"
	"net/url"
)

type Issue struct {
	Id             int64          `json:"id"`
	ProjectId      *int64         `json:"projectId"`
	IssueKey       string         `json:"issueKey"`
	KeyId          int64          `json:"keyId"`
	IssueType      *IssueType     `json:"issueType"`
	Summary        string         `json:"summary"`
	Description    string         `json:"description"`
	Resolution     *Resolution    `json:"resolution"`
	Priority       *Priority      `json:"priority"`
	Status         *ProjectStatus `json:"status"`
	Assignee       *User          `json:"assignee"`
	Category       []*Category    `json:"category"`
	Versions       []*Version     `json:"versions"`
	Milestone      []*Milestone   `json:"milestone"`
	StartDate      *Date          `json:"startDate"`
	DueDate        *Date          `json:"dueDate"`
	EstimatedHours *Hours         `json:"estimatedHours"`
	ActualHours    *Hours         `json:"actualHours"`
	ParentIssueId  *int64         `json:"parentIssueId"`
	CreatedUser    *User          `json:"createdUser"`
	Created        *Date          `json:"created"`
	UpdatedUser    *User          `json:"updatedUser"`
	Updated        *Date          `json:"updated"`
	CustomFields   []interface{}  `json:"customFields"`
	Attachments    []Attachment   `json:"attachments"`
	SharedFiles    []*SharedFile  `json:"sharedFiles"`
	Stars          []*Star        `json:"stars"`
}

func (i *Issue) URLValues() url.Values {
	query := url.Values{}

	if i == nil {
		return query
	}
	if i.ProjectId != nil {
		query.Add("projectId", fmt.Sprint(*i.ProjectId))
	}
	if i.Summary != "" {
		query.Add("summary", i.Summary)
	}
	if i.ParentIssueId != nil {
		query.Add("parentIssueId", fmt.Sprint(*i.ParentIssueId))
	}
	if i.Description != "" {
		query.Add("description", i.Description)
	}
	if i.Status != nil {
		query.Add("statusId", fmt.Sprint(i.Status.Id))
	}
	if i.Resolution != nil {
		query.Add("resolutionId", fmt.Sprint(i.Resolution.Id))
	}
	if i.StartDate != nil {
		query.Add("startDate", i.StartDate.String())
	}
	if i.DueDate != nil {
		query.Add("dueDate", i.DueDate.String())
	}
	if i.EstimatedHours != nil {
		query.Add("estimatedHours", fmt.Sprint(*i.EstimatedHours))
	}
	if i.ActualHours != nil {
		query.Add("actualHours", fmt.Sprint(*i.ActualHours))
	}
	if i.IssueType != nil {
		query.Add("issueTypeId", fmt.Sprint(i.IssueType.Id))
	}
	if i.Priority != nil {
		query.Add("priorityId", fmt.Sprint(i.Priority.Id))
	}
	if i.Assignee != nil {
		query.Add("assigneeId", fmt.Sprint(i.Assignee.Id))
	}
	if len(i.Category) > 0 {
		for _, category := range i.Category {
			query.Add("categoryId", fmt.Sprint(category.Id))
		}
	}
	if len(i.Versions) > 0 {
		for _, version := range i.Versions {
			query.Add("versionId", fmt.Sprint(version.Id))
		}
	}
	if len(i.Milestone) > 0 {
		for _, milestone := range i.Milestone {
			query.Add("milestoneId", fmt.Sprint(milestone.Id))
		}
	}
	if len(i.Attachments) > 0 {
		for _, attachment := range i.Attachments {
			query.Add("attachmentId", fmt.Sprint(attachment.Id))
		}
	}

	return query
}

func (a *Issue) Uniq(b *Issue) {
	if a.Summary == b.Summary {
		a.Summary = ""
	}
	if a.Description == b.Description {
		a.Description = ""
	}
	if a.IssueType != nil && b.IssueType != nil {
		if a.IssueType.Id == b.IssueType.Id {
			a.IssueType = nil
		}
	}
	if a.Resolution != nil && b.Resolution != nil {
		if a.Resolution.Id == b.Resolution.Id {
			a.Resolution = nil
		}
	}
	if a.Priority != nil && b.Priority != nil {
		if a.Priority.Id == b.Priority.Id {
			a.Priority = nil
		}
	}
	if a.Status != nil && b.Status != nil {
		if a.Status.Id == b.Status.Id {
			a.Status = nil
		}
	}
	if a.Assignee != nil && b.Assignee != nil {
		if a.Assignee.Id == b.Assignee.Id {
			a.Assignee = nil
		}
	}
	if a.StartDate != nil && b.StartDate != nil {
		if a.StartDate.Time().Equal(b.StartDate.Time()) {
			a.StartDate = nil
		}
	}
	if a.DueDate != nil && b.DueDate != nil {
		if a.DueDate.Time().Equal(b.DueDate.Time()) {
			a.DueDate = nil
		}
	}
	if a.EstimatedHours != nil && b.EstimatedHours != nil {
		if *a.EstimatedHours == *b.EstimatedHours {
			a.EstimatedHours = nil
		}
	}
	if a.ActualHours != nil && b.ActualHours != nil {
		if *a.ActualHours == *b.ActualHours {
			a.ActualHours = nil
		}
	}
}
