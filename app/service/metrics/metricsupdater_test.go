package metrics

import (
	"testing"

	"github.com/jo-hoe/routriever/app"
	"github.com/jo-hoe/routriever/app/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

type mockRouteService struct {
	metricValue int
	hitCount    int
}

func (m *mockRouteService) GetRouteDistance(start config.GPSCoordinates, end config.GPSCoordinates) (int, error) {
	m.hitCount++
	return m.metricValue, nil
}

func TestMetricsUpdater_updateMetric(t *testing.T) {
	testConfig := app.MetricConfig{
		Route: config.Route{
			Name: "Route to TomTom HQs",
			Start: config.GPSCoordinates{
				Latitude:  1.0,
				Longitude: 2.0,
			}, End: config.GPSCoordinates{
				Latitude:  3.0,
				Longitude: 4.0,
			},
		},
		Metric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "test_device",
			Help: "test device",
		}),
	}
	expectedValue := 10
	mockRouteService := mockRouteService{
		metricValue: expectedValue,
	}
	service := NewMetricsUpdater([]app.MetricConfig{testConfig}, &mockRouteService, 0)

	service.updateMetric(service.metricConfigs[0])

	if testutil.ToFloat64(testConfig.Metric) != float64(expectedValue) {
		t.Errorf("expected value %v, got %v", expectedValue, testutil.ToFloat64(testConfig.Metric))
	}
	if mockRouteService.hitCount != 1 {
		t.Errorf("expected hit count 1, got %v", mockRouteService.hitCount)
	}
}

func TestMetricsUpdater_NewMetricsUpdater(t *testing.T) {
	testConfig := app.MetricConfig{
		Route: config.Route{
			Name: "Route to TomTom HQs",
			Start: config.GPSCoordinates{
				Latitude:  1.0,
				Longitude: 2.0,
			}, End: config.GPSCoordinates{
				Latitude:  3.0,
				Longitude: 4.0,
			},
		},
		Metric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "test_device",
			Help: "test device",
		}),
	}

	updater := NewMetricsUpdater([]app.MetricConfig{testConfig}, &mockRouteService{}, 0)

	// check initial values
	expectedValue := -1.0
	if testutil.ToFloat64(updater.metricConfigs[0].Metric) != expectedValue {
		t.Errorf("expected value %v, got %v", expectedValue, testutil.ToFloat64(testConfig.Metric))
	}
}
