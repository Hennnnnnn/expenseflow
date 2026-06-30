package handlers

import (
	"net/http"
	"strconv"

	"github.com/Hennnnnnn/expenseflow/internal/email"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/response"
)

type EmailHandler struct {
	syncService *email.SyncService
}

func NewEmailHandler(syncService *email.SyncService) *EmailHandler {
	return &EmailHandler{
		syncService: syncService,
	}
}

func (h *EmailHandler) Sync(
	w http.ResponseWriter,
	r *http.Request,
) {

	limit := 100

	if value := r.URL.Query().Get("limit"); value != "" {

		n, err := strconv.Atoi(value)
		if err != nil {
			response.Error(
				w,
				http.StatusBadRequest,
				"Invalid limit parameter",
			)
			return
		}

		limit = n
	}

	result, err := h.syncService.SyncLatest(limit)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"Email synchronization completed successfully",
		result,
	)
}
