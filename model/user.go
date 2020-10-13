package model

import (
	"log"
	"time"
)

// User struct
type User struct {
	ID        uint      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// GetUserByID user id get
func GetUserByID(id int) (user User, err error) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}
