package payment

import (
	"errors"
	"fmt"
	log "tg_cs/logger"

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
		log.ErrorLogger.Fatal("Bad yookassa")
	}
	return nil
}

func CreatePayment(price int, description string) (string, string, error) {
	log.InfoLogger.Printf("(CreatePayment) price %d description %s", price, description)

	value := fmt.Sprintf("%d.00", price)

	paymentHandler := yookassa.NewPaymentHandler(yooClient)
	payment, err := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    value,
			Currency: "RUB",
		},
		Capture:       true,
		PaymentMethod: yoopayment.PaymentMethodType("sbp"),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "https://t.me/csbotwess_bot",
		},
		Description: description,
	})

	if err != nil {
		log.ErrorLogger.Fatalf("(CreatePayment) %v", err)
		return "", "", err
	}

	link, err := paymentHandler.ParsePaymentLink(payment)

	if err != nil {
		log.ErrorLogger.Fatalf("(ParsePaymentLink) %v", err)
		return "", "", err
	}

	payid := payment.ID

	fmt.Println(link)

	return link, payid, nil
}

func IsPaymentSuccess(orderId string) bool {
	status, err := GetPayment(orderId)
	if err != nil {
		fmt.Println("Error checking payment status:", err)
	} else {
		if status == "succeeded" {
			return true
		}
	}

	return false
}

func GetPayment(orderId string) (string, error) {
	log.InfoLogger.Printf("(GetPayment) orderId %s", orderId)
	paymentHandler := yookassa.NewPaymentHandler(yooClient)
	payment, _ := paymentHandler.FindPayment(orderId)
	if payment == nil {
		fmt.Println("payment not found")
		return "", errors.New("payment not found")
	}

	log.InfoLogger.Printf("(GetPayment) FindPayment [id: %s Status: %s]", payment.ID, string(payment.Status))

	return string(payment.Status), nil
}
