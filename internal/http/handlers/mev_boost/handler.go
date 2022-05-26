package mev_boost

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
)

type Handler struct {
	mevUC mev_boost.Usecase
}

func New(mevUC mev_boost.Usecase) *Handler {
	return &Handler{
		mevUC: mevUC,
	}
}

func (h *Handler) Handler(c *gin.Context) {
	var In struct {
		Message string `json:"message"`
		Nick    string `json:"nick"`
	}

	if parseErr := c.BindJSON(&In); parseErr != nil {
		c.JSON(http.StatusBadRequest, `Dad request`)

		return
	}
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": In.Message,
		"nick":    In.Nick,
	})
}
