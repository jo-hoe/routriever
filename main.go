package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/jo-hoe/routriever/app/service"
	"github.com/jo-hoe/routriever/app/service/gpsservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	portEnvVar  = "PORT"
	defaultPort = "8080"

	configPathEnvVar  = "CONFIG_PATH"
	defaultConfigPath = "./config.yaml"

	updateRateEnvVar  = "UPDATE_RATE"
	defaultUpdateRate = 1200
)

var (
	writer *service.ScheduledWriter
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", probeHandler)

	port := os.Getenv(portEnvVar)
	if port == "" {
		port = defaultPort
	}

	go writer.Run()

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func init() {
	configPath := os.Getenv(configPathEnvVar)
	if configPath == "" {
		log.Printf("CONFIG_PATH not set, using default path '%s'", defaultConfigPath)
		configPath = defaultConfigPath
	}

	serviceConfig, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal("could not read config")
	}

	gpsServiceInstance, err := gpsservice.NewRoutrieverService()
	if err != nil {
		log.Fatal("could not create routriever service")
	}

	updateRate := os.Getenv(updateRateEnvVar)
	duration, err := time.ParseDuration(updateRate)
	if err != nil {
		log.Printf("could not read duration %v, using default value %vs", err, defaultUpdateRate)
		duration = time.Duration(time.Duration(defaultUpdateRate) * time.Second)
	}

	writer = service.NewScheduledWriter(duration, &gpsServiceInstance, &serviceConfig)
}

func probeHandler(ctx echo.Context) (err error) {
	return ctx.NoContent(http.StatusOK)
}
