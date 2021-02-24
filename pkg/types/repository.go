package types

// Repository represents the git repository.
type Repository struct {
	Id           int64   `json:"id"`
	ProjectId    int64   `json:"projectId"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	HookURL      *string `json:"hookUrl"`
	HTTPURL      string  `json:"httpUrl"`
	SSHURL       string  `json:"sshUrl"`
	DisplayOrder int     `json:"displayOrder"`
	PushedAt     *Date   `json:"pushedAt"`
	CreatedUser  *User   `json:"createdUser"`
	Created      *Date   `json:"created"`
	UpdatedUser  *User   `json:"createdUser"`
	Updated      *Date   `json:"created"`
}
