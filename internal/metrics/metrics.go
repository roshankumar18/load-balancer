package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "loadbalancer_requests_total",
			Help: "Total number of requests served.",
		}, []string{"backend"},
	)

	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "loadbalancer_active_connections",
			Help: "Number of active connections per backend",
		}, []string{"backend"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "loadbalancer_request_duration_seconds",
			Help:    "Histogram of request durations per backend",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"backend"},
	)
)

func Init() {
	prometheus.MustRegister(RequestsTotal, ActiveConnections, RequestDuration)
}
