package model

// UserLayer User ORM layer
type UserLayer interface {
	GetUserByID(int) (User, error)
	CreateUser(User) (User, error)
}
