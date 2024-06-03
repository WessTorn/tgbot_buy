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

/*

// Выполнение SQL запроса
		rows, err := db.Query("SELECT hostname FROM amx_serverinfo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Обработка результатов запроса
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				log.Fatal(err)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, name)
			_, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}


		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
*/
