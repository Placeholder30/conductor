package user

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserStore interface {
	GetUserByEmail(email string)
	GetUserByID(id int)
	CreateUser(id string)
}
