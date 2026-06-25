package handlers

import (
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/email"
)

type EmailHandler struct {
	service *email.Service
}

func NewEmailHandler(service *email.Service) *EmailHandler {
	return &EmailHandler{
		service: service,
	}
}

func (h *EmailHandler) Sync(w http.ResponseWriter, r *http.Request) {

	err := h.service.SyncLatest(20)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Email sync completed"))
}
