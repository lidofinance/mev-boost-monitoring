package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lidofinance/mev-boost-monitoring/internal/transport/http/middlewares/rate_limiter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ulule/limiter/v3/drivers/store/memory"

	"github.com/lidofinance/mev-boost-monitoring/internal/transport/http/handlers/health"
	"github.com/lidofinance/mev-boost-monitoring/internal/transport/http/handlers/mev_boost"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func (app *App) RegisterRoutes(router *gin.Engine) {
	RateLimiterStore := memory.NewStore()
	rateMiddleware := rate_limiter.New(RateLimiterStore)

	router.Use(gin.LoggerWithWriter(app.Logger.Out))
	router.Use(gin.Recovery())
	router.Use(mgin.NewMiddleware(rateMiddleware))

	router.GET("/health", health.New().Handler)
	router.GET("/metrics", app.prometheusHandler())

	mevBoostHandler := mev_boost.New(app.Logger, app.usecase.MevBoost)

	router.POST("/payload", mevBoostHandler.HandlerPost)
	router.GET("/payload/:current_page/:per_page", mevBoostHandler.HandlerGet)
}

func (app *App) prometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(app.Metrics.Prometheus, promhttp.HandlerOpts{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
