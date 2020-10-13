package model

import (
	"database/sql"

	// mysql package
	_ "github.com/go-sql-driver/mysql"
)

// Connect Database Connect
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/laravel8?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, err
}
