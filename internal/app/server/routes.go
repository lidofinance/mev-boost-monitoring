package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/lidofinance/go-template/internal/http/handlers/health"
	userexample "github.com/lidofinance/go-template/internal/http/handlers/user_example"
)

func (app *App) RegisterRoutes(router *mux.Router) {
	handlers.RecoveryHandler()(router)

	router.HandleFunc("/health", health.New().Handler).Methods(http.MethodGet)
	router.Handle("/metrics", promhttp.HandlerFor(app.Metrics.Prometheus, promhttp.HandlerOpts{})).Methods(http.MethodGet)

	router.HandleFunc("/example", userexample.New(app.usecase.User).Handler).Methods(http.MethodGet)
}
