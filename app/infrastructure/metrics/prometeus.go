package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var httpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_requests_count",
	Help:        "The total number of requests received",
	ConstLabels: map[string]string{"version": "1.0.0"},
}, []string{"route", "method", "status_code"})

var httpRequestDurationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_request_duration_seconds_hist",
	Help:        "Time (in seconds) spent serving HTTP requests",
	ConstLabels: map[string]string{"version": "1.0.0"},
	Buckets:     prometheus.DefBuckets,
}, []string{"route", "method", "status_code"})

var httpRequestDurationSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_request_duration_seconds_summary",
	Help:        "Time (in seconds) spent serving HTTP requests",
	ConstLabels: map[string]string{"version": "1.0.0"},
	Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
}, []string{"route", "method", "status_code"})

var httpRequestDurationGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_request_duration_seconds_gauge",
	Help:        "The current number of items in the cache",
	ConstLabels: map[string]string{"version": "1.0.0"},
}, []string{"route", "method", "status_code"})

func MetricsRegister() {
	prometheus.MustRegister(httpRequestCounter)
	prometheus.MustRegister(httpRequestDurationHist)
	prometheus.MustRegister(httpRequestDurationSummary)
	prometheus.MustRegister(httpRequestDurationGauge)
}
