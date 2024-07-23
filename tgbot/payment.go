package tgbot

import (
	"database/sql"
	"tg_cs/database"
	"tg_cs/logger"
	"tg_cs/payment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowPayment(bot *tgbotapi.BotAPI, db *sql.DB, user *database.Context) {
	chatID := user.ChatID
	logger.Log.Debugf("(CreatePayment) User %d", chatID)

	// TODO: Обработать ощибки
	link, payid, _ := payment.CreatePayment()

	err := YooCreateMsg(bot, chatID, link)
	if err != nil {
		logger.Log.Fatalf("(YooCreateMsg) %v", err)
	}

	payment.AddPayData(chatID, payid, link)

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

func CancelPayment(chatID int64) {
	logger.Log.Debugf("(CancelPayment) User %d", chatID)

	payID, err := payment.GetPayIDFromChatID(chatID)
	if err != nil {
		logger.Log.Fatalf("(GetPayIDFromChatID) %v", err)
	}

	err = payment.CancelPayment(payID)
	if err != nil {
		logger.Log.Fatalf("(CancelPayment) %v", err)
	}
}
