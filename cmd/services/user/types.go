package user

type UserStore interface {
	GetUserByEmail(email string)
	GetUserByID(id int)
	CreateUser(id string)
}
