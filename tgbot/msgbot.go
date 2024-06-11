package tgbot

import (
	"database/sql"
	"fmt"
	"tg_cs/get_data"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WelcomeMsg(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Привет!")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.Send(msg)
	return err
}

func ServersMsg(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите сервер:")
	msg.ReplyMarkup = GetTeamButtons(bot, db)
	_, err := bot.Send(msg)
	return err
}

func ServiceMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите услугу:")
	msg.ReplyMarkup = GetServicesButtons(bot)
	_, err := bot.Send(msg)
	return err
}

func PrivilegesMsg(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите привилегию:")
	msg.ReplyMarkup = GetPrivilegesButton(bot)
	_, err := bot.Send(msg)
	return err
}

func PrivilegesDaysMsg(bot *tgbotapi.BotAPI, chatID int64, privilege get_data.Privilege) error {
	text := fmt.Sprintf("Вы выбрали услугу: %s.\nОписание:\n %s\nПожалуйста, выберите на сколько хотите взять привилегию:", privilege.Name, privilege.Description)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = GetPrivilegesDaysButton(bot, privilege)
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
