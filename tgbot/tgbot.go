package tgbot

import (
	"database/sql"
	"tg_cs/config"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTGBot() (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(config.BotToken())
}

func PlayTGBot(bot *tgbotapi.BotAPI, db *sql.DB) {
	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			logger.Log.Infof("(BOT) User %d set message: %v", chatID, update.Message.Text)

			if update.Message.Text == "/start" {
				ShowServersWelcome(bot, db, chatID)
				continue
			}

			var user *database.Context
			user, err := database.CtxGetUserData(db, chatID)
			if err != nil {
				logger.Log.Debugf("(CtxGetUserData) ERROR: %v", err)
				ShowServersWelcome(bot, db, chatID)
				continue
			}

			if user.ServiceID.Int64 == 1 {
				user.Privilege, err = database.CtxGetUserPrvgData(db, chatID)
				if err != nil {
					// TODO: Обработать если пропали данные
					logger.Log.Debugf("(CtxGetUserPrvgData) ERROR: %v", err)
					ShowPrivileges(bot, db, chatID)
					continue
				}
			}

			switch update.Message.Text {
			case "Вернуться назад":
				BackButton(bot, db, update, user)
				continue
			}

			switch user.Stage {
			case database.ServerStg:
				HandlerServers(bot, db, update, user)
			case database.ServiceStg:
				HandlerService(bot, db, update, user)
			case database.PrivilegeStg:
				HandlerPrivileges(bot, db, update, user)
			case database.PrvgDaysStg:
				HandlerPrivilegesDays(bot, db, update, user)
			case database.PrlgSteamStg:
				HandlerSteam(bot, db, update)
			case database.PrlgNickStg:
				HandlerNick(bot, db, update, user)
			}

		}
	}
}

func BackButton(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(BackButton) ChatID %d, User: %v", update.Message.Chat.ID, user)
	chatID := update.Message.Chat.ID
	err := database.UpdateBacking(db, user)
	if err != nil {
		logger.Log.Fatalf("(UpdateBacking) %v", err)
	}
	switch user.Stage {
	case database.ServiceStg:
		ShowServers(bot, db, chatID)
	case database.PrivilegeStg:
		ShowService(bot, db, chatID)
	case database.PrvgDaysStg:
		ShowPrivileges(bot, db, chatID)
	case database.PrlgSteamStg:
		ShowPrivilegesDays(bot, db, chatID, user.Privilege.PrvgID.Int64)
	case database.PrlgNickStg:
		ShowSteam(bot, db, chatID)
	}
}
