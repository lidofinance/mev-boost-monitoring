package health

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}

	jsonResponse, _ := json.Marshal(resp{Status: `ok`})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
