package mev_boost

import (
	"github.com/flashbots/go-boost-utils/types"
	"github.com/gin-gonic/gin"

	"github.com/bxcodec/faker/v3"
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

func (h *Handler) HandlerPost(c *gin.Context) {
	var In types.GetHeaderResponse

	_ = faker.FakeData(&In)

	err := h.mevUC.Create(c, In)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
	}
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func (h *Handler) HandlerGet(c *gin.Context) {
	out, err := h.mevUC.Get(c)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
	}

	c.JSON(200, out)
}
