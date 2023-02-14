package api

import (
	"observability-series-golang-edition/app/infrastructure/circuitbreaker"
	"observability-series-golang-edition/app/infrastructure/metrics"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRegister(e *echo.Echo, cb *circuitbreaker.CircuitBreaker) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(metrics.PrometheusMiddleware)
	e.Use(circuitbreaker.CircuitBreakerMiddleware(cb))
}
