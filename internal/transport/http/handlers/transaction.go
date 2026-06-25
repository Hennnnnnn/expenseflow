package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Hennnnnnn/expenseflow/internal/service"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/dto"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/response"
	"time"
	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"github.com/Hennnnnnn/expenseflow/internal/domain/types"
	"github.com/Hennnnnnn/expenseflow/internal/mapper"
	"github.com/google/uuid"
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

	transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"Invalid transaction date",
		)

		return
	}

	transaction := entity.Transaction{
		UUID:            uuid.NewString(),
		TransactionDate: transactionDate,
		Amount:          req.Amount,
		Type:            types.TransactionType(req.Type),
		Merchant:        req.Merchant,
		Description:     req.Description,
		Category:        req.Category,
		Source:          types.TransactionSource(req.Source),
		ReferenceNo:     req.ReferenceNo,
	}

	if err := h.service.Create(&transaction); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"transaction created successfully",
		mapper.ToTransactionResponse(transaction),
	)
}

func (h *TransactionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetAll()

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}	

	response.Success(
		w,
		http.StatusOK,
		"Transaction retrieved successfully",
		mapper.ToTransactionResponses(transactions),
	)
}