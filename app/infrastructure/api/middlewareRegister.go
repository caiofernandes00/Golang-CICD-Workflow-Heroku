package api

import (
	"overengineering-my-application/app/infrastructure/cache"
	"overengineering-my-application/app/infrastructure/circuitbreaker"
	"overengineering-my-application/app/infrastructure/metrics"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRegister(e *echo.Echo, cb *circuitbreaker.CircuitBreaker) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(cache.CacheMiddleware)
	e.Use(metrics.PrometheusMiddleware)
	e.Use(circuitbreaker.CircuitBreakerMiddleware(cb))
}
