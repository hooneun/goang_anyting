package model

import (
	"log"

	"github.com/hooneun/golang_anyting/helper"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName User table name
func (User) TableName() string {
	return "users"
}

// GetUserByID user id get
func (db *ORM) GetUserByID(id int) (user User, err error) {
	return user, db.Unscoped().First(&user, id).Error
}

// CreateUser add user!
func (db *ORM) CreateUser(u User) (user User, err error) {
	u.Password, err = helper.MakeHash(u.Password)
	if err != nil {
		return user, err
	}

	return user, db.Unscoped().Create(&user).Error
}

// LoginUser !
func (db *ORM) LoginUser(email, password string) bool {
	return false
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
