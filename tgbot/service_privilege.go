package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/game"
	"tg_cs/get_data"
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
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerPrivileges) User %d", chatID)
	privilegeName := update.Message.Text

	privilege, err := get_data.GetPrivilegeFromName(privilegeName)
	if err != nil {
		ShowPrivileges(bot, db, chatID)
		return
	}

	err = database.CtxInitUserPrvg(db, chatID, privilege.ID)
	if err != nil {
		logger.Log.Fatalf("(CtxInitUserPrvg) %v", err)
	}

	//ShowSteam(bot, db, chatID)
	ShowPrivilegesDays(bot, db, chatID, privilege.ID)
}

func ShowPrivilegesDays(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64, privilegeID int64) {
	logger.Log.Debugf("(ShowPrivilegesDays) User %d", chatID)

	privilege, err := get_data.GetPrivilegeFromID(privilegeID)
	if err != nil {
		ShowPrivileges(bot, db, chatID)
		return
	}

	err = PrivilegesDaysMsg(bot, chatID, privilege)
	if err != nil {
		logger.Log.Fatalf("(PrivilegesDaysMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrvgDaysStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerPrivilegesDays(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerPrivilegesDays) User %d (%v)", chatID, user)
	text := update.Message.Text

	privilege, err := get_data.GetPrivilegeFromID(user.Prvg.PrivilegeID.Int64)
	if err != nil {
		logger.Log.Debugf("(GetPrivilegeFromID) User %v", err)
		ShowPrivileges(bot, db, chatID)
		return
	}

	costID, err := get_data.GetCostIDFromString(privilege, text)
	if err != nil {
		logger.Log.Debugf("(GetCostIDFromString) User %v", err)
		ShowPrivilegesDays(bot, db, chatID, user.Prvg.PrivilegeID.Int64)
		return
	}

	err = database.CtxUpdateUserPrvgCostID(db, chatID, costID)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserPrvgCostID) %v", err)
	}

	ShowSteam(bot, db, chatID)
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

func HandlerSteam(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerSteam) User %d", chatID)
	steamID := update.Message.Text

	if !game.IsSteamIDValid(steamID) {
		ShowSteam(bot, db, chatID)
		return
	}

	err := database.CtxUpdateUserPrvgSteamID(db, chatID, steamID)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	ShowNick(bot, db, chatID)

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
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerNick) User %d", chatID)
	nick := update.Message.Text

	err := database.CtxUpdateUserPrvgNick(db, chatID, nick)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	user.Prvg.Nick.String = update.Message.Text

	database.SetAdminServer(db, user)

	err = PrivilegeMsg(bot, user.Prvg.ChatID)
	if err != nil {
		logger.Log.Fatalf("(PrivilegeMsg) %v", err)
	}

}
