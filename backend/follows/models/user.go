package models

type User struct {
	Result []struct{
		UserId			string		`json:"user_id"`
		OutFollowing	[]string	`json:"out_Following"`
		RId				string		`json:"@rid"`
	}	`json:"result"`
}