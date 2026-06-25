package handlers

import (
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/response"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, map[string]string {
		"status": "OK",
		"service": "ExpenseFlow",
		"version": "0.1.0",
	})
}