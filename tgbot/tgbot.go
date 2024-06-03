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
				ShowServers(bot, db, update.Message.Chat.ID)
				continue
			}

			var user *database.Context
			user, err := database.CtxGetUserData(db, update.Message.Chat.ID)
			if err != nil {
				logger.Log.Debugf("(CtxGetUserData) ERROR: %v", err)
				ShowServers(bot, db, update.Message.Chat.ID)
				continue
			}

			switch user.Stage {
			case database.PrivilegeStg:
				HandlerSteam(bot, db, update, user)
			case database.PrlgNickStg:
				HandlerNick(bot, db, update, user)
			}

		} else if update.CallbackQuery != nil {
			var user *database.Context
			user, err := database.CtxGetUserData(db, update.CallbackQuery.Message.Chat.ID)
			if err != nil {
				logger.Log.Debugf("(CtxGetUserData) ERROR: %v", err)
				ShowServers(bot, db, update.CallbackQuery.Message.Chat.ID)
				continue
			}

			switch user.Stage {
			case database.ServerStg:
				HandlerServers(bot, db, update, user)
			case database.ServiceStg:
				HandlerService(bot, db, update, user)
			case database.PrivilegeStg:
				HandlerPrivileges(bot, db, update, user)
			}

		}

	}
}
