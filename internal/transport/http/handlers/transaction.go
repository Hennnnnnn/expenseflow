package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Hennnnnnn/expenseflow/internal/service"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/dto"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/response"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler {
		service: s,
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"request parsed successfully",
		req,
	)
}

func (h *TransactionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	
}