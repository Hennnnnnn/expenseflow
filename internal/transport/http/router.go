package http

import (
	"log/slog"
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/transport/http/handlers"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	logger *slog.Logger,
	transactionHandler *handlers.TransactionHandler,
	emailHandler *handlers.EmailHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger(logger))

	r.Use(middleware.Recovery(logger))

	r.Get("/health", handlers.Health)
	r.Post("/transactions", transactionHandler.Create)
	r.Get("/transactions", transactionHandler.GetAll)
	r.Post("/api/email/sync", emailHandler.Sync)
	return r

}
