package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ORM database ORM struct
type ORM struct {
	*gorm.DB
}

// NewORM !
func NewORM() (*ORM, error) {
	dns := "root@tcp(127.0.0.1)/laravel8?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return &ORM{
		DB: db,
	}, err
}
