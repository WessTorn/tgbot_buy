package tgbot

import (
	"database/sql"
	"tg_cs/database"
	log "tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowServersWelcome(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	log.InfoLogger.Printf("(ShowServersWelcome) User %d", chatID)
	err := WelcomeMsg(bot, db, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(WelcomeMsg) %v", err)
	}

	err = database.CtxInitUser(db, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxInitUser) %v", err)
	}

	ShowServers(bot, db, chatID)
}

func ShowServers(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	log.InfoLogger.Printf("(ShowServers) User %d", chatID)
	err := ServersMsg(bot, db, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(ServersMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.ServerStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerServers(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	log.InfoLogger.Printf("(HandlerServersMenu) User %d", update.Message.Chat.ID)

	chatID := update.Message.Chat.ID
	serverName := update.Message.Text

	server, err := database.GetServerFromName(db, serverName)
	if err != nil {
		if err.Error() == "ServerNotFound" {
			err := BadButtonMsg(bot, db, user)
			if err != nil {
				log.ErrorLogger.Fatalf("(BadButtonMsg) %v", err)
			}
		}
		return
	}

	user.ChatID = chatID
	user.ServerID.Int64 = server.ID

	err = database.CtxUpdateUserServer(db, user)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateUserServer) %v", err)
	}

	ShowService(bot, db, chatID)
}
