package metrics

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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

var httpRequestCache = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_requests_cache",
	Help:        "The current number of items in the cache",
	ConstLabels: map[string]string{"version": "1.0.0"},
}, []string{"route", "method", "status_code"})

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}

		route := c.Path()
		if route == "" {
			route = "/"
		}
		method := c.Request().Method
		statusCode := strconv.Itoa(c.Response().Status)

		httpRequestCounter.WithLabelValues(route, method, fmt.Sprint(statusCode)).Inc()
		httpRequestDurationHist.WithLabelValues(route, method, fmt.Sprint(statusCode)).Observe(time.Since(start).Seconds())
		httpRequestDurationSummary.WithLabelValues(route, method, fmt.Sprint(statusCode)).Observe(time.Since(start).Seconds())
		httpRequestCache.WithLabelValues(route, method, fmt.Sprint(statusCode)).SetToCurrentTime()
		return nil
	}
}

func RegisterMetrics() {
	prometheus.MustRegister(httpRequestCounter)
	prometheus.MustRegister(httpRequestDurationHist)
	prometheus.MustRegister(httpRequestDurationSummary)
	prometheus.MustRegister(httpRequestCache)
}
