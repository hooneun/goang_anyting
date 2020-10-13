package model

import (
	"log"
	"time"
)

// User struct
type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// GetUserByID user id get
func GetUserByID(id int64) (user User, err error) {
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

// CreateUser add user!
func CreateUser(user User) (u User, err error) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}

	// if result.RowsAffected() < 1 {
	// 	return user
	// }

	lastID, _ := result.LastInsertId()
	return GetUserByID(lastID)
}

// DestroyUserByID delete user
func DestroyUserByID(id int64) (n int64, err error) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return result.RowsAffected()
}
