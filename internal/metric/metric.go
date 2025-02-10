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
	requestCounter prometheus.Counter
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
	}

	return nil
}

func IncRequestCounter() {
	metrics.requestCounter.Inc()
}
