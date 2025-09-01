package main

import (
	"tg_cs/config"
	"tg_cs/database"
	"tg_cs/get_data"
	"tg_cs/logger"
	log "tg_cs/logger"
	"tg_cs/payment"
	"tg_cs/tgbot"
)

func main() {
	config.InitConfig()

	payment.InitYookassaClient()

	logger.Init("./log/")

	get_data.ReadPrivilege()

	db, err := database.ConnectDB()
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}
	defer db.Close()

	log.InfoLogger.Println("Database connected")

	err = database.PingDB(db)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	log.InfoLogger.Println("Database ping successful")

	err = database.СtxCreate(db)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	err = database.СtxPrvgCreate(db)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	log.InfoLogger.Println("Context created")

	bot, err := tgbot.InitTGBot()
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	tgbot.PlayTGBot(bot, db)
}
