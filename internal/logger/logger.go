package logger

import (
	"log"
	"marketplace/internal/config"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger(cfg *config.Config) error {
	tempLogger := logrus.New()

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}
	tempLogger.SetLevel(level)

	Logger = tempLogger
	return nil
}
