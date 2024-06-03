package tgbot

import (
	"database/sql"
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
				showStartMenu(bot, db, update)
				continue
			}
		} else if update.CallbackQuery != nil {
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

func showStartMenu(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update) {
	logger.Log.Debugf("(showStartMenu) User %d", update.Message.Chat.ID)
	err := WecomeMsg(bot, db, update)
	if err != nil {
		logger.Log.Fatalf("(WecomeMsg) %v", err)
	}

	// err = data.InitUser(db, update.Message.Chat.ID)
	// if err != nil {
	// 	logger.Log.Fatalf("(InitUser) %v", err)
	// }
}
