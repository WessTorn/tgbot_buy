package logger

import (
	"tg_cs/config"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	level, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		logrus.Fatal(err)
	}

	Log.SetLevel(level)
}
