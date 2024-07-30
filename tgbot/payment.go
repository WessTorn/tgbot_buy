package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/get_data"
	"tg_cs/logger"
	"tg_cs/payment"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowPayment(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context) {
	chatID := user.ChatID
	logger.Log.Debugf("(CreatePayment) User %d", chatID)

	privelege, err := get_data.GetPrivilegeFromID(user.Privilege.PrvgID.Int64)
	if err != nil {
		logger.Log.Fatalf("(GetPrivilegeFromID) %v", err)
	}

	day := get_data.GetDayFromDayID(privelege, user.Privilege.DayID.Int64)

	// TODO: Обработать ощибки
	link, payid, _ := payment.CreatePayment(day.Price, privelege.Name)

	err = YooCreateMsg(bot, chatID, link)
	if err != nil {
		logger.Log.Fatalf("(YooCreateMsg) %v", err)
	}

	payment.AddPayData(chatID, payid, link)

	go checkPaymentStatus(bot, db, user, payid)

	err = database.CtxUpdateStage(db, chatID, database.PayYooStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}

func HandlerPayment(bot *tgbotapi.BotAPI, db *sql.DB, update tgbotapi.Update, user *database.Context) {
	chatID := update.Message.Chat.ID
	logger.Log.Debugf("(HandlerPayment) User %d", chatID)

	payID, err := payment.GetPayIDFromChatID(chatID)
	if err != nil {
		logger.Log.Fatalf("(GetPayIDFromChatID) %v", err)
	}

	status, err := payment.GetPayment(payID)

	err = YooStatusMsg(bot, chatID, status)
	if err != nil {
		logger.Log.Fatalf("(YooStatusMsg) %v", err)
	}
}

func checkPaymentStatus(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context, orderId string) {
	chatID := user.ChatID
	logger.Log.Debugf("(checkPaymentStatus) chatID %d", chatID)
	for {
		logger.Log.Debugf("(checkPaymentStatus) FOR %s", orderId)

		if payment.IsPaymentSuccess(orderId) {
			showSuccessPayment(bot, db, user)
			break
		}

		// Ждем 5 секунд перед следующей проверкой
		time.Sleep(5 * time.Second)
	}
}

func showSuccessPayment(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context) {
	chatID := user.ChatID
	logger.Log.Debugf("(showSuccessPayment) User %d", chatID)

	logger.Log.Debugf("(showSuccessPayment) User ServiceID id %d", user.ServiceID.Int64)

	err := YooSucceedMsg(bot, chatID)
	if err != nil {
		logger.Log.Fatalf("(YooSucceedMsg) %v", err)
	}

	switch user.ServiceID.Int64 {
	case 1:
		ShowFinishPrivilege(bot, db, user)
	}
}
