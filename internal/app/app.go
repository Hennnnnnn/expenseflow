package app

import {
	"log/slog"

	"github.com/Hennnnnnn/expenseflow/internal/config"
}

type App struct {
	Config *config.Config
	Logger *slog.Logger
}