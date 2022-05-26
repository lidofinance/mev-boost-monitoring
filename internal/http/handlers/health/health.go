package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handler(c *gin.Context) {
	type resp struct {
		Status string `json:"status"`
	}

	c.JSON(http.StatusOK, resp{Status: `ok`})
}
