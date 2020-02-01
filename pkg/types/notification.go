package types

type Notification struct {
	Id                  uint64      `json:"id"`
	AlreadyRead         bool        `json:"alreadyRead"`
	Reason              int         `json:"reason"`
	ResourceAlreadyRead bool        `json:"resourceAlreadyRead"`
	Project             Project     `json:"project"`
	Issue               Issue       `json:"issue"`
	Comment             Comment     `json:"comment"`
	PullRequest         PullRequest `json:"pullRequest"`
	PullRequestComment  Comment     `json:"pullRequestComment"`
	Sender              User        `json:"sender"`
	Created             Date        `json:"created"`
}
