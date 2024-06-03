package tgbot

import (
	"database/sql"
	"strconv"
	"tg_cs/database"
	"tg_cs/get_data"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetTeamButtons(bot *tgbotapi.BotAPI, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	servers := database.GetServers(db)

	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, server := range servers {
		serverID := strconv.Itoa(server.ID)
		button := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(server.Name, serverID))
		buttons = append(buttons, button)
	}

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func GetServiceButtons(bot *tgbotapi.BotAPI) tgbotapi.InlineKeyboardMarkup {
	var buttons [][]tgbotapi.InlineKeyboardButton
	button := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Покупка привилегии", "1"))
	buttons = append(buttons, button)
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func GetPrivilegesButton(bot *tgbotapi.BotAPI) tgbotapi.InlineKeyboardMarkup {
	privileges := get_data.GetPrivileges()

	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, privilege := range privileges.Privilege {
		prlgID := strconv.Itoa(privilege.ID)
		button := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(privilege.Name, prlgID))
		buttons = append(buttons, button)
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func GetBackButton(bot *tgbotapi.BotAPI) tgbotapi.ReplyKeyboardMarkup {
	button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Вернуться назад"))

	return tgbotapi.NewReplyKeyboard(button)
}
