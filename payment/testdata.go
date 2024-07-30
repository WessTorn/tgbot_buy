package payment

import "errors"

type TestPayData struct {
	UserID int64
	PayID  string
	Link   string
}

var TestPaysData []TestPayData

func AddPayData(chatID int64, payID string, link string) {
	var tempPay TestPayData
	tempPay.UserID = chatID
	tempPay.PayID = payID
	tempPay.Link = link

	TestPaysData = append(TestPaysData, tempPay)
}

func GetPayIDFromChatID(chatID int64) (string, error) {
	for _, pay := range TestPaysData {
		if pay.UserID != chatID {
			continue
		}

		return pay.PayID, nil
	}

	return "", errors.New("PayNotFound")
}
