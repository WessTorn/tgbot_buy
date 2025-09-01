package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/game"
	"tg_cs/get_data"
	log "tg_cs/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowPrivileges(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	log.InfoLogger.Printf("(ShowPrivileges) User %d", chatID)

	err := PrivilegesMsg(bot, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(PrivilegesMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrivilegeStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerPrivileges(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	log.InfoLogger.Printf("(HandlerPrivileges) User %d", chatID)
	privilegeName := update.Message.Text

	privilege, err := get_data.GetPrivilegeFromName(privilegeName)
	if err != nil {
		if err.Error() == "PrivilegeNotFound" {
			err := BadButtonMsg(bot, db, user)
			if err != nil {
				log.ErrorLogger.Fatalf("(BadButtonMsg) %v", err)
			}
		}
		return
	}

	err = database.CtxUpdateUserPrvgID(db, chatID, privilege.ID)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxInitUserPrvg) %v", err)
	}

	ShowPrivilegesDays(bot, db, chatID, privilege.ID)
}

func ShowPrivilegesDays(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64, privilegeID int64) {
	log.InfoLogger.Printf("(ShowPrivilegesDays) User %d", chatID)

	privilege, err := get_data.GetPrivilegeFromID(privilegeID)
	if err != nil {
		ShowPrivileges(bot, db, chatID)
		return
	}

	err = PrivilegesDaysMsg(bot, chatID, privilege)
	if err != nil {
		log.ErrorLogger.Fatalf("(PrivilegesDaysMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrvgDaysStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerPrivilegesDays(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	log.InfoLogger.Printf("(HandlerPrivilegesDays) User %d (%v)", chatID, user)
	text := update.Message.Text

	privilege, err := get_data.GetPrivilegeFromID(user.Privilege.PrvgID.Int64)
	if err != nil {
		log.InfoLogger.Printf("(GetPrivilegeFromID) User %v", err)
		ShowPrivileges(bot, db, chatID)
		return
	}

	dayID, err := get_data.GetDayIDFromString(privilege, text)
	if err != nil {
		// DayIDNotFound
		log.InfoLogger.Printf("(GetDayIDFromString) User %v", err)
		ShowPrivilegesDays(bot, db, chatID, user.Privilege.PrvgID.Int64)
		return
	}

	err = database.CtxUpdateUserPrvgDayID(db, chatID, dayID)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateUserPrvgDayID) %v", err)
	}

	ShowSteam(bot, db, chatID)
}

func ShowSteam(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	log.InfoLogger.Printf("(ShowSteam) User %d", chatID)
	err := SteamIDMsg(bot, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(SteamIDMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrlgSteamStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerSteam(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	log.InfoLogger.Printf("(HandlerSteam) User %d", chatID)
	steamID := update.Message.Text

	if !game.IsSteamIDValid(steamID) {
		ShowSteam(bot, db, chatID)
		return
	}

	err := database.CtxUpdateUserPrvgSteamID(db, chatID, steamID)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	ShowNick(bot, db, chatID)

}

func ShowNick(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64) {
	log.InfoLogger.Printf("(ShowNick) User %d", chatID)

	err := NickMsg(bot, chatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(NickMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrlgNickStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerNick(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	log.InfoLogger.Printf("(HandlerNick) User %d", chatID)
	nick := update.Message.Text

	err := database.CtxUpdateUserPrvgNick(db, chatID, nick)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateUserSteamID) %v", err)
	}

	user.Privilege.Nick.String = update.Message.Text

	ShowVerification(bot, db, user)
}

func ShowVerification(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context) {
	chatID := user.ChatID
	log.InfoLogger.Printf("(ShowVerification) User %d", chatID)

	err := VerificationMsg(bot, db, user)
	if err != nil {
		log.ErrorLogger.Fatalf("(VerificationMsg) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PrlgVerifStg)
	if err != nil {
		log.ErrorLogger.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerVerification(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	log.InfoLogger.Printf("(HandlerVerification) User %d", chatID)

	verification := update.Message.Text

	if verification != "Оплатить" {
		ShowVerification(bot, db, user)
	}

	ShowPayment(bot, db, user)

}

func ShowFinishPrivilege(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context) {
	//database.SetAdminServer(db, user)

	//TODO: Нормально все завершить.

	err := PrivilegeMsg(bot, user.Privilege.ChatID)
	if err != nil {
		log.ErrorLogger.Fatalf("(PrivilegeMsg) %v", err)
	}
}
