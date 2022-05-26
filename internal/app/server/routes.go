package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/lidofinance/mev-boost-monitoring/internal/http/handlers/health"
	"github.com/lidofinance/mev-boost-monitoring/internal/http/handlers/mev_boost"
)

func (app *App) RegisterRoutes(router *gin.Engine) {
	router.Use(gin.LoggerWithWriter(log.Writer()))
	router.Use(gin.Recovery())

	router.GET("/health", health.New().Handler)
	router.GET("/metrics", app.prometheusHandler())

	router.POST("/payload", mev_boost.New(app.usecase.MevBoost).HandlerPost)
	router.GET("/payload", mev_boost.New(app.usecase.MevBoost).HandlerGet)
}

func (app *App) prometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(app.Metrics.Prometheus, promhttp.HandlerOpts{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
