package service

import (
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

}
