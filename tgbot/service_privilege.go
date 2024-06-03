package tgbot

import (
	"database/sql"
	"strconv"
	"tg_cs/database"
	"tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowPrivileges(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowPrivileges) User %d", chatID)

	err := PrivilegesMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(PrivilegesMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrivilegeStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerPrivileges(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(HandlerPrivileges) User %d", update.CallbackQuery.Message.Chat.ID)

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	_, err := bot.Request(callback)
	if err != nil {
		logger.Log.Fatalf("(Request) %v", err)
	}

	privilegeID, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		logger.Log.Fatalf("(Atoi) %v", err)
	}

	err = database.CtxUpdateUserPrivilege(db, update.CallbackQuery.Message.Chat.ID, privilegeID)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	ShowSteam(bot, db, update.CallbackQuery.Message.Chat.ID)

}

func ShowSteam(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowSteam) User %d", chatID)
	err := SteamIDMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(SteamIDMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrlgSteamStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerSteam(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("(HandlerSteam) User %d", update.Message.Chat.ID)

	err := database.CtxUpdateUserSteamID(db, update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	ShowNick(bot, db, update.Message.Chat.ID)

}

func ShowNick(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	logger.Log.Debugf("(ShowNick) User %d", chatID)
	err := NickMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(NickMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrlgNickStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerNick(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	logger.Log.Debugf("%v", user)
	logger.Log.Debugf("(HandlerNick) User %d", update.Message.Chat.ID)

	err := database.CtxUpdateUserNick(db, update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	user.Nick.String = update.Message.Text

	database.SetAdminServer(db, user)

	err = PrivilegeMsg(bot, user.ChatID)
	if err != nil {
		logger.Log.Fatalf("(PrivilegeMsg) %v", err)
	}

}
