package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		logrus.Fatal(err)
	}

	Log.SetLevel(level)
}
