package models

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
