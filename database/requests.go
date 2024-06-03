package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetServersData(db *sql.DB) []Server {
	var servers []Server
	rows, err := db.Query("SELECT id, hostname, address FROM amx_serverinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var server Server
		err = rows.Scan(&server.ID, &server.Name, &server.IP)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}

	return servers
}
