package app

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/prometheus/client_golang/prometheus"
)

var prometheusMetricNamePattern = regexp.MustCompile("[a-zA-Z_:][a-zA-Z0-9_:]*")

func GeneratePrometheusMetrics(config config.Config) map[string]prometheus.Gauge {
	metrics := make(map[string]prometheus.Gauge)

	for _, route := range config.Routes {
		metricNameSlice := prometheusMetricNamePattern.FindAllString(route.Name, -1)
		metricName := strings.Join(metricNameSlice, "_")

		metric := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: metricName,
			Help: fmt.Sprintf("Distance in second driving for route %s", route.Name),
		})
		metrics[route.Name] = metric
	}

	return metrics
}
