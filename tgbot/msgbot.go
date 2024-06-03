package tgbot

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ServersMsg(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Привет!\nВыберите сервер:")
	msg.ReplyMarkup = GetTeamButtons(bot, db)
	_, err := bot.Send(msg)
	return err
}
