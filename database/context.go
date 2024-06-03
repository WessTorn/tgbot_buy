package database

import (
	"database/sql"
	"fmt"
	"tg_cs/logger"
)

func СtxCreate(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tgbot_ctx (
			id INT(11) NOT NULL auto_increment PRIMARY KEY,
            chat_id INT(11) NULL DEFAULT NULL,
			stage INT(11) NULL DEFAULT NULL,
			server_id INT(11) NULL DEFAULT NULL,
			service INT(11) NULL DEFAULT NULL,
			privilege_id INT(11) NULL DEFAULT NULL,
			steam_id VARCHAR(24) NULL DEFAULT NULL,
			nick VARCHAR(24) NULL DEFAULT NULL
        )
    `)
	return err
}

func CtxInitUser(db *sql.DB, chatID int64) error {
	logger.Log.Debug("(CtxInitUser)")
	var count int

	sqlReq := "SELECT COUNT(*) FROM tgbot_ctx WHERE chat_id = ?"
	err := db.QueryRow(sqlReq, chatID).Scan(&count)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	if count > 0 {
		sqlReq = "UPDATE tgbot_ctx SET stage = ?, server_id = NULL WHERE chat_id = ?"
		_, err = db.Exec(sqlReq, ServerStg, chatID)
		if err != nil {
			return fmt.Errorf("(%s): %v", sqlReq, err)
		}
	} else {
		sqlReq = "INSERT INTO tgbot_ctx (chat_id, stage) VALUES (?, ?)"
		_, err = db.Exec(sqlReq,
			chatID,
			ServerStg,
		)
		if err != nil {
			return fmt.Errorf("(%s): %v", sqlReq, err)
		}
	}

	return nil
}

func CtxUpdateStage(db *sql.DB, chatID int64, stage Stage) error {
	logger.Log.Debug("(CtxUpdateStage)")

	sqlReq := "UPDATE tgbot_ctx SET stage = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, stage, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxGetUserData(db *sql.DB, chatID int64) (*Context, error) {
	logger.Log.Debug("(CtxGetUserData)")
	var user Context

	row := db.QueryRow("SELECT chat_id, stage, server_id, steam_id, nick FROM tgbot_ctx WHERE chat_id = ?", chatID)

	err := row.Scan(&user.ChatID, &user.Stage, &user.ServerID, &user.SteamID, &user.Nick)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with chat_id %d not found", chatID)
		}
		return nil, err
	}

	return &user, nil
}

func CtxUpdateUserServer(db *sql.DB, chatID int64, serverID int) error {
	logger.Log.Debug("(CtxUpdateUserServer)")

	sqlReq := "UPDATE tgbot_ctx SET server_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, serverID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserService(db *sql.DB, chatID int64, serverID int) error {
	logger.Log.Debug("(CtxUpdateUserService)")

	sqlReq := "UPDATE tgbot_ctx SET service = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, serverID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserPrivilege(db *sql.DB, chatID int64, privilegeID int) error {
	logger.Log.Debug("(CtxUpdateUserPrivilege)")

	sqlReq := "UPDATE tgbot_ctx SET privilege_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, privilegeID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserSteamID(db *sql.DB, chatID int64, steamID string) error {
	logger.Log.Debug("(CtxUpdateUserSteamID)")

	sqlReq := "UPDATE tgbot_ctx SET steam_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, steamID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserNick(db *sql.DB, chatID int64, nick string) error {
	logger.Log.Debug("(CtxUpdateUserNick)")

	sqlReq := "UPDATE tgbot_ctx SET nick = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, nick, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}
