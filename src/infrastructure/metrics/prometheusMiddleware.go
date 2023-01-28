package metrics

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

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
		httpRequestDurationGauge.WithLabelValues(route, method, fmt.Sprint(statusCode)).Set(time.Since(start).Seconds())
		return nil
	}
}
