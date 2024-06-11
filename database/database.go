package database

import (
	"database/sql"
	"fmt"
	"tg_cs/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	text := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DBUser(), config.DBPass(), config.DBHost(), config.DBDatabase())
	return sql.Open("mysql", text)
}

func PingDB(db *sql.DB) error {
	return db.Ping()
}
