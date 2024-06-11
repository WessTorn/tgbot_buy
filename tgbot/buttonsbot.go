package tgbot

import (
	"database/sql"
	"fmt"
	"tg_cs/database"
	"tg_cs/get_data"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetTeamButtons(bot *tgbotapi.BotAPI, db *sql.DB) tgbotapi.ReplyKeyboardMarkup {
	servers := database.GetServers(db)

	var buttons [][]tgbotapi.KeyboardButton
	for _, server := range servers {
		button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(server.Name))
		buttons = append(buttons, button)
	}

	return tgbotapi.NewReplyKeyboard(buttons...)
}

func GetServicesButtons(bot *tgbotapi.BotAPI) tgbotapi.ReplyKeyboardMarkup {
	services := get_data.GetServices()

	var buttons [][]tgbotapi.KeyboardButton

	for _, service := range services {
		button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(service))
		buttons = append(buttons, button)
	}

	buttonBack := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Вернуться назад"))
	buttons = append(buttons, buttonBack)

	return tgbotapi.NewReplyKeyboard(buttons...)
}

func GetPrivilegesButton(bot *tgbotapi.BotAPI) tgbotapi.ReplyKeyboardMarkup {
	privileges := get_data.GetPrivileges()

	var buttons [][]tgbotapi.KeyboardButton
	for _, privilege := range privileges.Privilege {
		button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(privilege.Name))
		buttons = append(buttons, button)
	}

	buttonBack := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Вернуться назад"))
	buttons = append(buttons, buttonBack)

	return tgbotapi.NewReplyKeyboard(buttons...)
}

func GetPrivilegesDaysButton(bot *tgbotapi.BotAPI, privilege get_data.Privilege) tgbotapi.ReplyKeyboardMarkup {
	var buttons [][]tgbotapi.KeyboardButton
	for _, cost := range privilege.Cost {
		text := fmt.Sprintf("%d дней - %d руб", cost.Day, cost.Price)
		button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(text))
		buttons = append(buttons, button)
	}

	buttonBack := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Вернуться назад"))
	buttons = append(buttons, buttonBack)

	return tgbotapi.NewReplyKeyboard(buttons...)
}

func GetBackButton(bot *tgbotapi.BotAPI) tgbotapi.ReplyKeyboardMarkup {
	button := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Вернуться назад"))

	return tgbotapi.NewReplyKeyboard(button)
}
