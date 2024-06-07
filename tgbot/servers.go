package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowServersWelcome(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowServersWelcome) User %d", chatID)
	err := WelcomeMsg(bot, db, chatID)
	if err != nil {
		logger.Log.Fatalf("(WelcomeMsg) %v", err)
	}

	err = database.CtxInitUser(db, chatID)
	if err != nil {
		logger.Log.Fatalf("(CtxInitUser) %v", err)
	}

	ShowServers(bot, db, chatID)
}

func ShowServers(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowServers) User %d", chatID)
	err := ServersMsg(bot, db, chatID)
	if err != nil {
		logger.Log.Fatalf("(ServersMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.ServerStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerServers(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(HandlerServersMenu) User %d", update.Message.Chat.ID)

	chatID := update.Message.Chat.ID
	serverName := update.Message.Text

	server, err := database.GetServerFromName(db, serverName)
	if err != nil {
		ShowServers(bot, db, chatID)
		return
	}

	user.ChatID = chatID
	user.ServerID.Int64 = server.ID

	err = database.CtxUpdateUserServer(db, user)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserServer) %v", err)
	}

	ShowService(bot, db, chatID)
}
