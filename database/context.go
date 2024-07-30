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
			service INT(11) NULL DEFAULT NULL
		);
	`)
	return err
}

func СtxPrvgCreate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tgbot_ctx_privilege (
			id INT(11) NOT NULL auto_increment PRIMARY KEY,
			chat_id INT(11) NULL DEFAULT NULL,
			privilege_id INT(11) NULL DEFAULT NULL,
			day_id INT(11) NULL DEFAULT NULL,
			steam_id VARCHAR(24) NULL DEFAULT NULL,
			nick VARCHAR(24) NULL DEFAULT NULL
		);
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
		sqlReq = "UPDATE tgbot_ctx SET stage = ?, server_id = NULL, service = NULL WHERE chat_id = ?"
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

	row := db.QueryRow("SELECT chat_id, stage, server_id, service FROM tgbot_ctx WHERE chat_id = ?", chatID)

	err := row.Scan(&user.ChatID, &user.Stage, &user.ServerID, &user.ServiceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with chat_id %d not found", chatID)
		}
		return nil, err
	}

	return &user, nil
}

func UpdateBacking(db *sql.DB, user *Context) error {
	var sqlReq string

	switch user.Stage {
	case ServiceStg:
		sqlReq = "UPDATE tgbot_ctx SET server_id = NULL WHERE chat_id = ?"
	case PrivilegeStg:
		sqlReq = "UPDATE tgbot_ctx SET service = NULL WHERE chat_id = ?"
	case PrvgDaysStg:
		sqlReq = "UPDATE tgbot_ctx_privilege SET privilege_id = NULL WHERE chat_id = ?"
	case PrlgSteamStg:
		sqlReq = "UPDATE tgbot_ctx_privilege SET day_id = NULL WHERE chat_id = ?"
	case PrlgNickStg:
		sqlReq = "UPDATE tgbot_ctx_privilege SET steam_id = NULL WHERE chat_id = ?"
	case PayYooStg:
		sqlReq = "UPDATE tgbot_ctx_privilege SET nick = NULL WHERE chat_id = ?"
	}

	_, err := db.Exec(sqlReq, user.ChatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserServer(db *sql.DB, user *Context) error {
	logger.Log.Debug("(CtxUpdateUserServer)")

	sqlReq := "UPDATE tgbot_ctx SET server_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, user.ServerID.Int64, user.ChatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserService(db *sql.DB, chatID int64, serviceID int64) error {
	logger.Log.Debug("(CtxUpdateUserService)")

	sqlReq := "UPDATE tgbot_ctx SET service = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, serviceID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxInitUserPrvg(db *sql.DB, chatID int64) error {
	logger.Log.Debug("(CtxInitUserPrvg)")
	var count int

	sqlReq := "SELECT COUNT(*) FROM tgbot_ctx_privilege WHERE chat_id = ?"
	err := db.QueryRow(sqlReq, chatID).Scan(&count)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	if count <= 0 {
		sqlReq = "INSERT INTO tgbot_ctx_privilege (chat_id) VALUES (?)"
		_, err = db.Exec(sqlReq,
			chatID,
		)
		if err != nil {
			return fmt.Errorf("(%s): %v", sqlReq, err)
		}
	}

	return nil
}

func CtxGetUserPrvgData(db *sql.DB, chatID int64) (ContextPrlg, error) {
	logger.Log.Debug("(CtxGetUserPrvgData)")
	var userPrlg ContextPrlg

	row := db.QueryRow("SELECT chat_id, privilege_id, day_id, steam_id, nick FROM tgbot_ctx_privilege WHERE chat_id = ?", chatID)

	err := row.Scan(&userPrlg.ChatID, &userPrlg.PrvgID, &userPrlg.DayID, &userPrlg.SteamID, &userPrlg.Nick)
	if err != nil {
		if err == sql.ErrNoRows {
			return userPrlg, fmt.Errorf("user with chat_id %d not found", chatID)
		}
		return userPrlg, err
	}

	return userPrlg, nil
}

func CtxUpdateUserPrvgID(db *sql.DB, chatID int64, privilegeID int64) error {
	logger.Log.Debug("(CtxUpdateUserPrvgID)")

	sqlReq := "UPDATE tgbot_ctx_privilege SET privilege_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, privilegeID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserPrvgDayID(db *sql.DB, chatID int64, dayID int64) error {
	logger.Log.Debug("(CtxUpdateUserPrvgDayID)")

	sqlReq := "UPDATE tgbot_ctx_privilege SET day_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, dayID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserPrvgSteamID(db *sql.DB, chatID int64, steamID string) error {
	logger.Log.Debug("(CtxUpdateUserPrvgSteamID)")

	sqlReq := "UPDATE tgbot_ctx_privilege SET steam_id = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, steamID, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}

func CtxUpdateUserPrvgNick(db *sql.DB, chatID int64, nick string) error {
	logger.Log.Debug("(CtxUpdateUserPrvgNick)")

	sqlReq := "UPDATE tgbot_ctx_privilege SET nick = ? WHERE chat_id = ?"
	_, err := db.Exec(sqlReq, nick, chatID)
	if err != nil {
		return fmt.Errorf("(%s): %v", sqlReq, err)
	}

	return nil
}
