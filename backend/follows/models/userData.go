package models

type UserData struct {
	Users []struct{
		Name	string		`json:"name"`
		Email	string		`json:"email"`
		Image	string		`json:"image"`
		ID		string		`json:"_id"`
	} `json:"users"`
}