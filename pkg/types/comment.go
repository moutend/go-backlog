package types

type Comment struct {
	Id            int64           `json:"id"`
	Content       string          `json:"content"`
	ChangeLog     []*ChangeLog    `json:"changeLog"`
	CreatedUser   *User           `json:"createdUser"`
	Created       *Date           `json:"created"`
	Updated       *Date           `json:"updated"`
	Stars         []*Star         `json:"stars"`
	Notifications []*Notification `json:"notifications"`
}
