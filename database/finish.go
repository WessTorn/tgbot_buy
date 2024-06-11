package database

import (
	"database/sql"
	"log"
	"tg_cs/game"
	"tg_cs/get_data"
	"tg_cs/logger"
	"time"
)

func SetAdminServer(db *sql.DB, user *Context) {
	privelege, err := get_data.GetPrivilegeFromID(user.Privilege.PrvgID.Int64)
	if err != nil {
		logger.Log.Fatalf("(GetPrivilegeFromID) %v", err)
	}

	nowTime := time.Now().Unix()
	day := get_data.GetDayFromCostID(privelege, user.Privilege.CostID.Int64)
	daysToAdd := day * 24 * 60 * 60
	futureTime := nowTime + int64(daysToAdd)

	sqlReq := "INSERT INTO amx_amxadmins (username, password, access, flags, steamid, nickname, ashow, created, expired, days ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(sqlReq,
		user.Privilege.Nick.String,
		"",
		privelege.Flags,
		"ce",
		user.Privilege.SteamID.String,
		user.Privilege.Nick.String,
		"1",
		nowTime,
		futureTime,
		day,
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

	server := GetServerFromId(db, int(user.ServerID.Int64))

	game.SendRCON(server.IP, server.Rcon, "amx_reloadadmins")
}

func GetAdminID(db *sql.DB, user *Context) int {
	var adminID int
	rows, err := db.Query("SELECT id FROM amx_amxadmins WHERE steamid = ?", user.Privilege.SteamID.String)
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
