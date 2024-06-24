package database

import (
	"database/sql"
	"errors"
	"log"
	"tg_cs/logger"

	_ "github.com/go-sql-driver/mysql"
)

func GetServers(db *sql.DB) []Server {
	var servers []Server
	sqlReq := "SELECT id, hostname, address FROM amx_serverinfo"
	rows, err := db.Query(sqlReq)
	if err != nil {
		logger.Log.Fatalf("(%s): %v", sqlReq, err)
	}
	defer rows.Close()

	for rows.Next() {
		var server Server
		err = rows.Scan(&server.ID, &server.Name, &server.IP)
		if err != nil {
			logger.Log.Fatalf("(GetServers - Scan): %v", err)
		}
		servers = append(servers, server)
	}

	return servers
}

func GetServerFromName(db *sql.DB, serverName string) (Server, error) {
	var server Server
	sqlReq := "SELECT id, hostname, address, rcon FROM amx_serverinfo WHERE hostname = ?"
	err := db.QueryRow(sqlReq, serverName).Scan(&server.ID, &server.Name, &server.IP, &server.Rcon)
	if err == sql.ErrNoRows {
		return server, errors.New("ServerNotFound")
	} else if err != nil {
		logger.Log.Fatalf("(%s): %v", sqlReq, err)
	}

	return server, nil
}

func GetServerFromId(db *sql.DB, serverID int) Server {
	var server Server
	rows, err := db.Query("SELECT id, hostname, address, rcon FROM amx_serverinfo WHERE id = ?", serverID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&server.ID, &server.Name, &server.IP, &server.Rcon)
		if err != nil {
			log.Fatal(err)
		}
	}

	return server
}
