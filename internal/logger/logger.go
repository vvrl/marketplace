package logger

import (
	"log"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

// func NewLogger() error {
// 	Logger = logrus.New()

// 	return nil
// }

func ConfigureLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	Logger.SetLevel(level)
	return nil
}
