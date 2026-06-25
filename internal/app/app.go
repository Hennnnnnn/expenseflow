package app

import (
	"log/slog"
	"gorm.io/gorm"
	"github.com/Hennnnnnn/expenseflow/internal/config"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *gorm.DB
}