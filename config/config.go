package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel   string
	DBHost     string
	DBUser     string
	DBPass     string
	DBDatabase string
	BotToken   string
}

var data Config

func InitConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("Failed to load .env file:", err)
	}

	data = Config{
		LogLevel:   os.Getenv("LOG_LEVEL"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPass:     os.Getenv("DB_PASS"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		BotToken:   os.Getenv("BOT_TOKEN"),
	}
}

func LogLevel() string {
	return data.LogLevel
}

func DBHost() string {
	return data.DBHost
}

func DBUser() string {
	return data.DBUser
}

func DBPass() string {
	return data.DBPass
}

func DBDatabase() string {
	return data.DBDatabase
}

func BotToken() string {
	return data.BotToken
}
