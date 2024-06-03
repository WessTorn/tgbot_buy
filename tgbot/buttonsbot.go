package tgbot

import (
	"database/sql"
	"strconv"
	"tg_cs/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetTeamButtons(bot *tgbotapi.BotAPI, db *sql.DB) tgbotapi.InlineKeyboardMarkup {
	servers := database.GetServersData(db)

	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, server := range servers {
		dataIP := strconv.Itoa(server.ID)
		button := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(server.Name, dataIP))
		buttons = append(buttons, button)
	}

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
