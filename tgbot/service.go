package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/get_data"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowService(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowService) User %d", chatID)
	err := ServiceMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(ServiceMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.ServiceStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerService(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerService) User %d", chatID)
	serviceName := update.Message.Text

	service, err := get_data.GetServiceFromName(serviceName)
	if err != nil {
		if err.Error() == "ServiceNotFound" {
			err := BadButtonMsg(bot, db, user)
			if err != nil {
				logger.Log.Fatalf("(BadButtonMsg) %v", err)
			}
		}
		return
	}

	err = database.CtxUpdateUserService(db, chatID, service.ID)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserService) %v", err)
	}

	switch service.ID {
	case 1:
		err = database.CtxInitUserPrvg(db, chatID)
		if err != nil {
			logger.Log.Fatalf("(CtxInitUserPrvg) %v", err)
		}
		ShowPrivileges(bot, db, chatID)
	}
}
