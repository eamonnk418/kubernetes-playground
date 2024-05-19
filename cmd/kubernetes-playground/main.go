package main

import (
	"context"
	"kubernetes-playground/internal/app"
	"kubernetes-playground/internal/config"
	"kubernetes-playground/internal/constants"
	"kubernetes-playground/internal/logger"
	"kubernetes-playground/pkg/cmd"
	"time"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.NewLogger()
	app := app.NewApp(cfg, logger)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, constants.AppCtxKey, app)

	cmd.Execute(ctx)
}
