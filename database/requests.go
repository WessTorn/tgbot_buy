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

func GetServerFromName(db *sql.DB, serverName string) (Server, error) {
	var server Server
	sqlReq := "SELECT id, hostname, address FROM amx_serverinfo WHERE hostname = ?"
	err := db.QueryRow(sqlReq, serverName).Scan(&server.ID, &server.Name, &server.IP)
	if err == sql.ErrNoRows {
		return server, errors.New("ServerNotFound")
	} else if err != nil {
		logger.Log.Fatalf("(%s): %v", sqlReq, err)
	}

	return server, nil
}

func GetServerFromId(db *sql.DB, serverID int) Server {
	var server Server
	rows, err := db.Query("SELECT id, hostname, address FROM amx_serverinfo WHERE id = ?", serverID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&server.ID, &server.Name, &server.IP)
		if err != nil {
			log.Fatal(err)
		}
	}

	return server
}

func SetAdminServer(db *sql.DB, user *Context) {
	sqlReq := "INSERT INTO amx_amxadmins (username, password, access, flags, steamid, nickname, ashow, expired, days ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlReq,
		user.Prvg.Nick.String,
		"",
		"abcdefghijklmnopqrstuv",
		"ce",
		user.Prvg.SteamID.String,
		user.Prvg.Nick.String,
		"1",
		0,
		0,
	)
	if err != nil {
		log.Fatalf("(%s): %v", sqlReq, err)
	}

	adminID := GetAdminID(db, user)

	sqlReq = "INSERT INTO amx_admins_servers (admin_id, server_id) VALUES (?, ?)"
	_, err = db.Exec(sqlReq,
		adminID,
		user.ServerID,
	)
	if err != nil {
		log.Fatalf("(%s): %v", sqlReq, err)
	}
}

func GetAdminID(db *sql.DB, user *Context) int {
	var adminID int
	rows, err := db.Query("SELECT id FROM amx_amxadmins WHERE steamid = ?", user.Prvg.SteamID.String)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&adminID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return adminID
}
