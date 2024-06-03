package tgbot

import (
	"database/sql"
	"strconv"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowService(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowService) User %d", chatID)
	err := ServiceMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(ServiceMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.ServiceStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerService(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(HandlerService) User %d", update.CallbackQuery.Message.Chat.ID)

	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	_, err := bot.Request(callback)
	if err != nil {
		logger.Log.Fatalf("(Request) %v", err)
	}

	service, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		logger.Log.Fatalf("(Atoi) %v", err)
	}

	err = database.CtxUpdateUserService(db, update.CallbackQuery.Message.Chat.ID, service)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserService) %v", err)
	}

	ShowPrivileges(bot, db, update.CallbackQuery.Message.Chat.ID)
}
