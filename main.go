package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jo-hoe/routriever/app/config"
	app "github.com/jo-hoe/routriever/app/service"
	"github.com/jo-hoe/routriever/app/service/gpsservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	portEnvVar  = "PORT"
	defaultPort = "8080"

	configPathEnvVar  = "CONFIG_PATH"
	defaultConfigPath = "./config.yaml"
)

var (
	metrics           map[string]prometheus.Gauge
	routrieverService gpsservice.RoutrieverService
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", probeHandler)
	http.Handle("/metrics", promhttp.Handler())

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

	config, err := config.GetConfigFromFile(configPath)
	if err != nil {
		log.Fatalf("could not read config %v", err)
	}

	// create metrics from config
	metrics = app.GeneratePrometheusMetrics(config)
	for _, metric := range metrics {
		prometheus.MustRegister(metric)
	}

	routrieverService, err = gpsservice.NewRoutrieverService()
	if err != nil {
		log.Fatal("could not create routriever service")
	}
}

func probeHandler(ctx echo.Context) (err error) {
	return ctx.NoContent(http.StatusOK)
}
