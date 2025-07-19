package app

import (
	"marketplace/internal/config"
	"marketplace/internal/handlers"
	"marketplace/internal/logger"
	"marketplace/internal/storage"

	"github.com/labstack/echo/v4"
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

	_, err := storage.NewDatabase(cfg)
	if err != nil {
		logger.Logger.Fatal(err)
	}
	logger.Logger.Info("database successfully started")

	e := echo.New()

	handlers.SetAPI(e)
	e.Logger.Fatal(e.Start(cfg.Server.Port))

	logger.Logger.Infof("server successfully started on %s port", cfg.Server.Port)

	return nil
}
