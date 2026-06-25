package handlers

import (
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/response"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {

	response.Success(w, http.StatusOK, "Service is healthy", map[string]string{
		"status":  "OK",
		"service": "ExpenseFlow",
		"version": "0.1.0",
	})
}