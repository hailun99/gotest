package dbutil

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
