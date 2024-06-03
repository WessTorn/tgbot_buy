package database

import (
	"database/sql"
	"fmt"
	"tg_cs/logger"
)

func СtxCreate(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tgbot_context (
			id INTEGER PRIMARY KEY,
            chat_id INTEGER,
			stage INTEGER,
			server_id INTEGER
        )
    `)
	return err
}

func CtxInitUser(db *sql.DB, chatID int64) error {
	logger.Log.Debug("(CtxInitUser)")
	var count int

	sqlReq := "SELECT COUNT(*) FROM tgbot_context WHERE chat_id = ?"
	err := db.QueryRow(sqlReq, chatID).Scan(&count)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	if count > 0 {
		sqlReq = "UPDATE tgbot_context SET stage = ?, server_id = NULL WHERE chat_id = ?"
		_, err = db.Exec(sqlReq, ServerStage, chatID)
		if err != nil {
			return fmt.Errorf("(%s): %v", sqlReq, err)
		}
	} else {
		sqlReq = "INSERT INTO tgbot_context (chat_id, stage, server_id) VALUES (?, ?, ?)"
		_, err = db.Exec(sqlReq,
			chatID,
			ServerStage,
			sql.NullInt64{},
		)
		if err != nil {
			return fmt.Errorf("(%s): %v", sqlReq, err)
		}
	}

	return nil
}

func CtxGetUserData(db *sql.DB, chatID int64) (*Context, error) {
	logger.Log.Debug("(CtxGetUserData)")
	var user Context

	row := db.QueryRow("SELECT chat_id, stage, server_id FROM tgbot_context WHERE chat_id = ?", chatID)

	err := row.Scan(&user.ChatID, &user.Stage, &user.ServerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with chat_id %d not found", chatID)
		}
		return nil, err
	}

	return &user, nil
}
