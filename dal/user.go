package dal

import ()

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (m User) Validate() string {
	if m.UserName == "" {
		return "Empty username"
	}
	if m.FirstName == "" {
		return "Empty first name"
	}
	if m.LastName == "" {
		return "Empty last name"
	}
	return ""
}
