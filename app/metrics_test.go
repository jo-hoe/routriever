package app

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestGeneratePrometheusMetrics(t *testing.T) {
	expectedMetricName := "Route between TomTom HQs"
	expectedMetricHelp := "Distance in second driving for route Route between TomTom HQs"

	type args struct {
		config config.Config
	}
	tests := []struct {
		name string
		args args
		want map[string]prometheus.Gauge
	}{
		{
			name: "test that metrics are generated",
			args: args{
				config: config.Config{
					Routes: []config.Route{
						{
							Name: "Route between TomTom HQs",
							Start: config.GPSCoordinates{
								Latitude:  1.0,
								Longitude: 2.0,
							},
							End: config.GPSCoordinates{
								Latitude:  3.0,
								Longitude: 4.0,
							},
						},
					},
				}},
			want: map[string]prometheus.Gauge{
				"Route between TomTom HQs": prometheus.NewGauge(prometheus.GaugeOpts{
					Name: expectedMetricName,
					Help: expectedMetricHelp,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GeneratePrometheusMetrics(tt.args.config)
			for key, value := range got {
				if _, ok := tt.want[key]; !ok {
					t.Errorf("GeneratePrometheusMetrics() = %v, want %v", got, tt.want)
				}

				if !strings.Contains(value.Desc().String(), expectedMetricName) {
					t.Errorf("GeneratePrometheusMetrics() = %v, want %v", got, tt.want)
				}

				if !strings.Contains(value.Desc().String(), expectedMetricHelp) {
					t.Errorf("GeneratePrometheusMetrics() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestRegisterMetrics(t *testing.T) {
	type args struct {
		metrics map[string]prometheus.Gauge
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test that metrics are registered",
			args: args{
				metrics: map[string]prometheus.Gauge{
					"test": prometheus.NewGauge(prometheus.GaugeOpts{
						Name: "test",
						Help: "test",
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterMetrics(tt.args.metrics)

			for _, metric := range tt.args.metrics {
				foundMetrics := testutil.CollectAndCount(metric)
				if foundMetrics != len(tt.args.metrics) {
					t.Errorf("RegisterMetrics() = %v, want %v", foundMetrics, len(tt.args.metrics))
				}
			}
		})
	}
}

func TestGetMetricsConfig(t *testing.T) {
	testRoute := config.Route{
		Name:  "Route between TomTom HQs",
		Start: config.GPSCoordinates{},
		End:   config.GPSCoordinates{},
	}
	testMetric := prometheus.NewGauge(prometheus.GaugeOpts{})

	type args struct {
		config  config.Config
		metrics map[string]prometheus.Gauge
	}
	tests := []struct {
		name string
		args args
		want []MetricConfig
	}{
		{
			name: "test that metrics and configs are combined correctly",
			args: args{
				config: config.Config{
					Routes: []config.Route{testRoute},
				},
				metrics: map[string]prometheus.Gauge{
					testRoute.Name: testMetric,
				},
			},
			want: []MetricConfig{
				{
					Route:  testRoute,
					Metric: testMetric,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMetricsConfig(tt.args.config, tt.args.metrics); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetricsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
