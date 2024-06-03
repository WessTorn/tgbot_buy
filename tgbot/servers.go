package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowServersMenu(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(showServerMenu) User %d", chatID)
	err := ServersMsg(bot, db, chatID)
	if err != nil {
		logger.Log.Fatalf("(ServersMsg) %v", err)
	}

	err = database.CtxInitUser(db, chatID)
	if err != nil {
		logger.Log.Fatalf("(CtxInitUser) %v", err)
	}
}
