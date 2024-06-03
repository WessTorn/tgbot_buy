package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// "username:password@tcp(127.0.0.1:3306)/dbname"

func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
}

func PingDB(db *sql.DB) error {
	return db.Ping()
}
