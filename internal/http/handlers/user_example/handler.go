package userexample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lidofinance/go-template/internal/pkg/users"
)

type Handler struct {
	userUc users.Usecase
}

func New(userUc users.Usecase) *Handler {
	return &Handler{
		userUc: userUc,
	}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user, err := h.userUc.Get(r.Context(), int64(1))
	if err != nil {
		fmt.Fprint(w, "user not found")
		return
	}

	jsonResponse, _ := json.Marshal(user)
	_, _ = w.Write(jsonResponse)
}
