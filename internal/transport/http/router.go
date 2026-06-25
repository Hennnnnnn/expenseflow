package http

import (
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/transport/http/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/health", handlers.Health)

	return r

}