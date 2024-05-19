package app

import (
	"kubernetes-playground/internal/config"
	"kubernetes-playground/internal/logger"
)

type App struct {
	*config.Config
	*logger.Logger
}

func NewApp(cfg *config.Config, logger *logger.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
	}
}
