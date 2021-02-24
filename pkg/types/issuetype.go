package types

type IssueType struct {
	Id           int64  `json:"id"`
	ProjectId    int64  `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}
