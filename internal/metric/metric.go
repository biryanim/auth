package metric

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "auth_space"
	appName   = "auth_app"
)

type Metrics struct {
	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_requests_total",
				Help:      "Количество запросов к серверу",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_response_total",
				Help:      "Количество ответов от сервера",
			},
			[]string{"status", "method"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_histogram_response_time",
				Help:      "Время ответа от сервера",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status"},
		),
	}

	return nil
}

func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

func IncResponseCounter(status, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

func HistogramResponseTimeObserve(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
