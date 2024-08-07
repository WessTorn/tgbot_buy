package main

import (
	"log"
	"tg_cs/config"
	"tg_cs/database"
	"tg_cs/get_data"
	"tg_cs/logger"
	"tg_cs/payment"
	"tg_cs/tgbot"
)

func main() {
	config.InitConfig()

	payment.InitYookassaClient()

	logger.InitLogger()

	get_data.ReadPrivilege()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger.Log.Info("Database connected")

	err = database.PingDB(db)
	if err != nil {
		log.Fatal(err)
	}

	logger.Log.Info("Database ping successful")

	err = database.СtxCreate(db)
	if err != nil {
		log.Fatal(err)
	}

	err = database.СtxPrvgCreate(db)
	if err != nil {
		log.Fatal(err)
	}

	logger.Log.Info("Context created")

	bot, err := tgbot.InitTGBot()
	if err != nil {
		log.Fatal(err)
	}

	tgbot.PlayTGBot(bot, db)
}
