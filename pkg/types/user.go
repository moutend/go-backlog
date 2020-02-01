package types

type User struct {
	Id           uint64       `json:"id"`
	UserId       string       `json:"userId"`
	Name         string       `json:"name"`
	RoleType     int          `json:"roleType"`
	Lang         string       `json:"lang"`
	MailAddress  string       `json:"mailAddress"`
	NulabAccount NulabAccount `json:"nulabAccount"`
}
