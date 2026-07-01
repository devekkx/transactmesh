package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests processed",
	},
	[]string{"service", "route", "status"},
)

func Register() {
	prometheus.MustRegister(HttpRequestsTotal)
}
