package metrics

import (
	"log"
	"time"

	"github.com/jo-hoe/routriever/app"
	"github.com/jo-hoe/routriever/app/service/gpsservice"
)

type MetricsUpdater struct {
	metricConfigs  []app.MetricConfig
	routeService   gpsservice.RoutrieverService
	updateInterval time.Duration
}

func NewMetricsUpdater(metricConfigs []app.MetricConfig, routeService gpsservice.RoutrieverService, updateInterval time.Duration) *MetricsUpdater {
	setInitialMetrics(metricConfigs)

	return &MetricsUpdater{
		metricConfigs:  metricConfigs,
		routeService:   routeService,
		updateInterval: updateInterval,
	}
}

func setInitialMetrics(metricConfigs []app.MetricConfig) {
	for _, metricConfig := range metricConfigs {
		metricConfig.Metric.Set(-1)
	}
}

func (m *MetricsUpdater) Start() {
	// create metrics initially
	m.updateMetrics()

	// schedule future metric updates
	// based on the update interval
	ticker := time.NewTicker(m.updateInterval)
	go func() {
		for range ticker.C {
			m.updateMetrics()
		}
	}()
}

func (m *MetricsUpdater) updateMetrics() {
	for _, metricConfig := range m.metricConfigs {
		go m.updateMetric(metricConfig)
	}
}

func (m *MetricsUpdater) updateMetric(metricConfig app.MetricConfig) {
	distance, err := m.routeService.GetRouteDistance(metricConfig.Route.Start, metricConfig.Route.End)
	if err != nil {
		log.Printf("could not get distance for route '%s' error: %v", metricConfig.Route.Name, err)
		return
	}
	metricConfig.Metric.Set(float64(distance))
	log.Printf("successfully updated metric for route '%s'", metricConfig.Route.Name)
}
