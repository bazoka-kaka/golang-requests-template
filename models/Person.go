package models

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json:"online"`
}
