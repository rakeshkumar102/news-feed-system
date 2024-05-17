package models

type User struct {
	Users	[]struct {
	ID		string		`json:"_id"`
	Name	string		`json:"name"`
	Image	string		`json:"image"`
	}	`json:"users"`
}