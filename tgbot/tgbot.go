package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTGBot() (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI("")
}

func PlayTGBot(bot *tgbotapi.BotAPI, db *sql.DB) {
	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				ShowServersMenu(bot, db, update.Message.Chat.ID)
				continue
			}
		} else if update.CallbackQuery != nil {
			var user *database.Context
			user, err := database.CtxGetUserData(db, update.CallbackQuery.Message.Chat.ID)
			if err != nil {
				logger.Log.Debugf("(CtxGetUserData) ERROR: %v", err)
				ShowServersMenu(bot, db, update.CallbackQuery.Message.Chat.ID)
				continue
			}

			switch user.Stage {
			case database.ServerStage:
				logger.Log.Debugf("%v", user)
				// Respond to the callback query, telling Telegram to show the user
				// a message with the data received.
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				// And finally, send a message containing the data received.
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}

		}

	}
}
