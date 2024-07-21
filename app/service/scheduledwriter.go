package service

import (
	"log"
	"time"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/jo-hoe/routriever/app/service/gpsservice"
)

type ScheduledWriter struct {
	interval time.Duration
	service  *gpsservice.RoutrieverService
	config   *config.Config
}

func NewScheduledWriter(interval time.Duration, service *gpsservice.RoutrieverService, config *config.Config) *ScheduledWriter {
	return &ScheduledWriter{
		interval: interval,
		service:  service,
		config:   config,
	}
}

func (sw *ScheduledWriter) Run() {
	for {
		go sw.processConfigs()
		time.Sleep(sw.interval)
	}
}

func (sw *ScheduledWriter) processConfigs() {
	for _, config := range sw.config.Routes {
		go processConfig(config, *sw.service)
	}
}

func processConfig(config config.Route, service gpsservice.RoutrieverService) {
	result, err := service.GetRouteDistance(config.Start, config.End)
	if err != nil {
		log.Printf("could not retrieve data for route %s error: %v", config.Name, err.Error())
	}
	log.Printf("retrieved data for route %s", config.Name)
	write(config, result)
}

func write(config config.Route, valueInSeconds int) {
	log.Printf("writing data for route %s with value %d", config.Name, valueInSeconds)
}
