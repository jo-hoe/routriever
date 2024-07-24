package app

import (
	"testing"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/prometheus/client_golang/prometheus"
)

func TestGeneratePrometheusMetrics(t *testing.T) {
	type args struct {
		config config.Config
	}
	tests := []struct {
		name string
		args args
		want map[string]prometheus.Gauge
	}{
		{
			name: "positive test",
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
					Name: "Route_between_TomTom_HQs",
					Help: "Distance in second driving for route Route between TomTom HQs",
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

				if value.Desc().String() != tt.want[key].Desc().String() {
					t.Errorf("GeneratePrometheusMetrics() = %v, want %v", value.Desc().String(), tt.want[key].Desc().String())
				}
			}
		})
	}
}
