package models

type User struct {
	UserID     string
	Role       *string
	Mail       *string
	Login      *string
	Password   *string
	TelegramID int64
}

func NewUser() *User {
	user := new(User)
	user.Role = new(string)
	user.Mail = new(string)
	user.Login = new(string)
	user.Password = new(string)
	return user
}
