package mev_boost

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
	"github.com/lidofinance/mev-boost-monitoring/internal/transport/http/dto"
	"github.com/lidofinance/mev-boost-monitoring/internal/utils/deps"
)

type Handler struct {
	log   deps.Logger
	mevUC mev_boost.Usecase
}

func New(log deps.Logger, mevUC mev_boost.Usecase) *Handler {
	return &Handler{
		log:   log,
		mevUC: mevUC,
	}
}

func (h *Handler) HandlerPost(c *gin.Context) {
	var req dto.CustomRelayPayload

	// TODO security
	// Rate-Limiter - done
	// Some Secret Key - for mvp - it's redundant
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(`Could not bind json`, err)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": `ShouldBindJSON err`,
			"err": err.Error(),
		})

		return
	}

	in, err := req.Validate()
	if err != nil {
		h.log.Error(`Request is invalid`, err)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": `Validation err`,
			"err": err.Error(),
		})

		return
	}

	if createErr := h.mevUC.Create(c, in); createErr != nil {
		h.log.Error(fmt.Errorf(`could not store payload %w`, createErr))

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": createErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "posted",
	})
}

func (h *Handler) HandlerGet(c *gin.Context) {
	currentPageIn := c.Param("current_page")
	perPageIn := c.Param("per_page")

	currentPage, err := strconv.ParseUint(currentPageIn, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": http.StatusText(http.StatusBadRequest),
		})

		return
	}

	if currentPage <= 0 {
		currentPage = 1
	}

	perPage, err := strconv.ParseUint(perPageIn, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": http.StatusText(http.StatusBadRequest),
		})

		return
	}

	if perPage <= 0 {
		perPage = 10
	}

	out, err := h.mevUC.Paginated(c, currentPage, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": http.StatusText(http.StatusInternalServerError),
		})
	}

	c.JSON(200, out)
}
