package domain

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	SaveUser(user User) error
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
}
