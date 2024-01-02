package entity

type User struct {
	ID             int64  `db:"id"`
	Email          string `db:"email"`
	Name           string `db:"name"`
	LastName       string `db:"lastname"`
	Password       string `db:"password"`
	Phone          string `db:"phone"`
	ProfilePicture string `db:"profile_picture"`
	Address        string `db:"address"`
}
