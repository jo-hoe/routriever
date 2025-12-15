package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jo-hoe/routriever/app"
	"github.com/jo-hoe/routriever/app/config"
	"github.com/jo-hoe/routriever/app/service/gpsservice"
	"github.com/jo-hoe/routriever/app/service/metrics"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	portEnvVar  = "PORT"
	defaultPort = "8080"

	configPathEnvVar  = "CONFIG_PATH"
	defaultConfigPath = "/config/config.yaml"

	secretPathEnvVar  = "SECRET_PATH"
	defaultSecretPath = "/secret/secret.txt"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", probeHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	port := os.Getenv(portEnvVar)
	if port == "" {
		port = defaultPort
	}

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func init() {
	configPath := os.Getenv(configPathEnvVar)
	if configPath == "" {
		log.Printf("CONFIG_PATH not set, using default path '%s'", defaultConfigPath)
		configPath = defaultConfigPath
	}

	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("could not read config %v", err)
	}

	interval, err := time.ParseDuration(config.UpdateInterval)
	if err != nil {
		log.Fatalf("could not parse update interval error: '%v'", err)
	}

	// this will panic in case config is incorrect
	prometheusMetrics := app.GeneratePrometheusMetrics(config)
	app.RegisterMetrics(prometheusMetrics)
	metricConfigs := app.GetMetricsConfig(config, prometheusMetrics)

	secretPath := os.Getenv(secretPathEnvVar)
	if secretPath == "" {
		log.Printf("SECRET_PATH not set, using default path '%s'", defaultSecretPath)
		secretPath = defaultSecretPath
	}
	routrieverService, err := gpsservice.NewRoutrieverService(secretPath)
	if err != nil {
		log.Fatalf("could not create routriever service error: '%v'", err)
	}

	metricsUpdater := metrics.NewMetricsUpdater(metricConfigs, routrieverService, interval)
	metricsUpdater.Start()
}

func probeHandler(ctx echo.Context) (err error) {
	return ctx.NoContent(http.StatusOK)
}
