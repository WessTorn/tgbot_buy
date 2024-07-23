package payment

import (
	"errors"
	"fmt"
	"tg_cs/logger"

	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

var yooClient *yookassa.Client

func InitYookassaClient() error {
	yooClient = yookassa.NewClient("407586", "test_tZs22tk326Jcqp2V3EQYHyx77cuXQeSVwAFaPnS8DNY")
	// Создаем обработчик настроек
	settingsHandler := yookassa.NewSettingsHandler(yooClient)
	// Получаем информацию о настройках магазина или шлюза
	settings, _ := settingsHandler.GetAccountSettings(nil)

	fmt.Println(settings)

	if settings == nil {
		logger.Log.Fatalf("Bad yookassa")
	}
	return nil
}

func CreatePayment() (string, string, error) {
	paymentHandler := yookassa.NewPaymentHandler(yooClient)
	payment, err := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    "1000.00",
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("sbp"),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "https://t.me/csbotwess_bot",
		},
		Description: "Test payment",
	})

	link, _ := paymentHandler.ParsePaymentLink(payment)
	payid := payment.ID

	fmt.Println(link)

	return link, payid, err
}

func CancelPayment(orderId string) error {
	paymentHandler := yookassa.NewPaymentHandler(yooClient)
	payment, _ := paymentHandler.FindPayment(orderId)
	if payment == nil {
		fmt.Println("payment not found")
		return errors.New("payment not found")
	}
	_, err := paymentHandler.CancelPayment(payment.ID)

	return err
}

func GetPayment(orderId string) (string, error) {
	paymentHandler := yookassa.NewPaymentHandler(yooClient)
	payment, _ := paymentHandler.FindPayment(orderId)
	if payment == nil {
		fmt.Println("payment not found")
		return "", errors.New("payment not found")
	}

	fmt.Printf("id: %s \nStatus: %s", payment.ID, string(payment.Status))

	return string(payment.Status), nil
}
