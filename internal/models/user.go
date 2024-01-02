package models

type User struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	LastName       string `json:"lastname"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
	Address        string `json:"address"`
}
