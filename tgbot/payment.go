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

	link, _ := payment.CreatePayment()

	msg := tgbotapi.NewMessage(chatID, link)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Fatalf("(ShowPayment) %v", err)
	}

	err = database.CtxUpdateStage(db, chatID, database.PayYooStg)
	if err != nil {
		logger.Log.Fatalf("(CtxUpdateStage) %v", err)
	}
}
