package tgbot

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WecomeMsg(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!\nВыберите сервер:")
	msg.ReplyMarkup = GetTeamButtons(bot, db)
	_, err := bot.Send(msg)
	return err
}
