package app

import (
	"marketplace/internal/config"
	"marketplace/internal/logger"
)

type App struct {
	config *config.Config
}

func NewApp() *App {
	return &App{
		config: config.NewConfig(),
	}
}

func (a *App) Run() error {

	cfg := config.NewConfig()
	logger.ConfigureLogger(cfg.Logger.Level)

	return nil
}
