package metrics

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "requests_total",
	Help: "The total number of requests received",
})

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestCounter.Inc()
		return next(c)
	}
}

func RegisterMetrics() {
	prometheus.MustRegister(requestCounter)
}
