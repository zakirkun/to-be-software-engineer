package router

import "github.com/prometheus/client_golang/prometheus"

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests processed, labeled by status and path.",
		},
		[]string{"status", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}
