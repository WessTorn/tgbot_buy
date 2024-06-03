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

func ServiceMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите услугу:")
	msg.ReplyMarkup = GetServiceButtons(bot)
	_, err := bot.Send(msg)
	return err
}

func PrivilegesMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите привилегию:")
	msg.ReplyMarkup = GetPrivilegesButton(bot)
	_, err := bot.Send(msg)
	return err
}

func SteamIDMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите ваш Steam ID:")
	msg.ReplyMarkup = GetBackButton(bot)
	_, err := bot.Send(msg)
	return err
}

func NickMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите ваш ник:")
	msg.ReplyMarkup = GetBackButton(bot)
	_, err := bot.Send(msg)
	return err
}

func PrivilegeMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Вам выдана привилегия")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.Send(msg)
	return err
}
