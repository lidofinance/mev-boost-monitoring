package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lidofinance/mev-boost-monitoring/internal/app/server"
	"github.com/lidofinance/mev-boost-monitoring/internal/connectors/logger"
	"github.com/lidofinance/mev-boost-monitoring/internal/connectors/metrics"
	"github.com/lidofinance/mev-boost-monitoring/internal/connectors/postgres"
	"github.com/lidofinance/mev-boost-monitoring/internal/env"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, envErr := env.Read(ctx)
	if envErr != nil {
		println("Read env error:", envErr.Error())

		os.Exit(1)
	}

	db, errDB := postgres.Connect(cfg.PgConfig)
	if errDB != nil {
		println("Connect db error:", errDB.Error())
		os.Exit(1)
	}

	log, logErr := logger.New(&cfg.AppConfig)

	if logErr != nil {
		println("Logger error:", logErr.Error())

		os.Exit(1)
	}

	log.Info(fmt.Sprintf(`started %s application`, cfg.AppConfig.Name))

	r := gin.Default()
	metrics := metrics.New(prometheus.NewRegistry(), cfg.AppConfig.Name, cfg.AppConfig.Env)

	repo := server.Repository(db)
	usecase := server.Usecase(repo)

	app := server.New(log, metrics, usecase, repo)

	app.Metrics.BuildInfo.Inc()
	app.RegisterRoutes(r)

	if err := server.RunHTTPServer(ctx, cfg.AppConfig.Port, r); err != nil {
		println(err)
	}
}
