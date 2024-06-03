package tgbot

import (
	"database/sql"
	"strconv"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowServers(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
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

func HandlerServers(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(HandlerServersMenu) User %d", update.CallbackQuery.Message.Chat.ID)

	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	_, err := bot.Request(callback)
	if err != nil {
		logger.Log.Fatalf("(Request) %v", err)
	}

	serverId, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		logger.Log.Fatalf("(Atoi) %v", err)
	}

	err = database.CtxUpdateUserServer(db, update.CallbackQuery.Message.Chat.ID, serverId)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserServer) %v", err)
	}

	ShowService(bot, db, update.CallbackQuery.Message.Chat.ID)

	// // And finally, send a message containing the data received.
	// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	// if _, err := bot.Send(msg); err != nil {
	// 	panic(err)
	// }
}
