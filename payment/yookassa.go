package payment

import (
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

func CreatePayment() (string, error) {
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

	fmt.Println(link)

	return link, err
}
