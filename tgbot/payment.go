package tgbot

import (
	"database/sql"
	"fmt"
	"tg_cs/database"
	"tg_cs/logger"
	"tg_cs/payment"
	"time"

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

	go checkPaymentStatus(bot, chatID, payid)

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

func checkPaymentStatus(bot *tgbotapi.BotAPI, chatID int64, orderId string) {
	logger.Log.Debugf("(checkPaymentStatus) orderId %s", orderId)
	for {
		logger.Log.Debugf("(checkPaymentStatus) FOR %s", orderId)
		// Проверяем статус платежа
		status, err := payment.GetPayment(orderId)
		if err != nil {
			fmt.Println("Error checking payment status:", err)
		} else {
			if status == "succeeded" {
				// Платеж успешно оплачен, выдаем привилегию
				err = PrivilegeMsg(bot, chatID)
				if err != nil {
					logger.Log.Fatalf("(PrivilegeMsg) %v", err)
				}
				break
			}
		}

		// Ждем 5 секунд перед следующей проверкой
		time.Sleep(5 * time.Second)
	}
}
