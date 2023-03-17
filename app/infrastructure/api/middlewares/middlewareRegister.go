package middlewares

import (
	"overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker"
	"overengineering-my-application/app/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRegister(
	e *echo.Echo, config *util.Config, cb *circuitbreaker.CircuitBreaker,
	loggerSetup middleware.RequestLoggerConfig, gzipSetup middleware.GzipConfig,
) {
	e.Use(middleware.RequestLoggerWithConfig(loggerSetup))
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(RateLimiterMiddleware(config)))
	e.Use(middleware.GzipWithConfig(gzipSetup))
	e.Use(CacheMiddleware(config))
	e.Use(PrometheusMiddleware)
	e.Use(CircuitBreakerMiddleware(cb))
}
