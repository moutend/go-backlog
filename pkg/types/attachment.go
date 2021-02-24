package types

type Attachment struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	CreatedUser *User  `json:"createdUser"`
	Created     *Date  `json:"created"`
	UpdatedUser *User  `json:"updatedUser"`
	Updated     *Date  `json:"updated"`
}
