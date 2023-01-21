package metrics

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_requests_count",
	Help:        "The total number of requests received",
	ConstLabels: map[string]string{"version": "1.0.0"},
})

var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   "echo_framework",
	Subsystem:   "rest_api",
	Name:        "http_request_duration_seconds",
	Help:        "Time (in seconds) spent serving HTTP requests",
	ConstLabels: map[string]string{"version": "1.0.0"},
}, []string{"route", "method", "code"})

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestCounter.Inc()

		route := c.Path()
		if route == "" {
			route = "/"
		}
		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}

		requestDuration.WithLabelValues(route).Observe(time.Since(start).Seconds())
		return nil
	}
}

func RegisterMetrics() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}
