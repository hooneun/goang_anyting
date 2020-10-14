package model

// UserLayer User ORM layer
type DBLayer interface {
	GetUserByID(int) (User, error)
	CreateUser(User) (User, error)
}
