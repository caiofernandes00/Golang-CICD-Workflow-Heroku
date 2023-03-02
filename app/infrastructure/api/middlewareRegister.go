package api

import (
	"overengineering-my-application/app/infrastructure/cache"
	"overengineering-my-application/app/infrastructure/circuitbreaker"
	"overengineering-my-application/app/infrastructure/metrics"
	"overengineering-my-application/app/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRegister(
	e *echo.Echo, config *util.Config, cb *circuitbreaker.CircuitBreaker,
	loggerSetup middleware.RequestLoggerConfig, gzipSetup middleware.GzipConfig) {

	e.Use(middleware.RequestLoggerWithConfig(loggerSetup))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(gzipSetup))
	e.Use(cache.CacheMiddleware(config))
	e.Use(metrics.PrometheusMiddleware)
	e.Use(circuitbreaker.CircuitBreakerMiddleware(cb))
}
