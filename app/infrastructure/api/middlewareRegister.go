package api

import (
	"overengineering-my-application/app/infrastructure/cache"
	"overengineering-my-application/app/infrastructure/circuitbreaker"
	"overengineering-my-application/app/infrastructure/metrics"
	"overengineering-my-application/app/util"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func MiddlewareRegister(e *echo.Echo, config *util.Config, cb *circuitbreaker.CircuitBreaker, logger zerolog.Logger) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			for _, s := range config.SkipCompressionUrls {
				if strings.Contains(c.Request().URL.Path, s) {
					return true
				}
			}
			return false
		},
	}))
	e.Use(cache.CacheMiddleware(config))
	e.Use(metrics.PrometheusMiddleware)
	e.Use(circuitbreaker.CircuitBreakerMiddleware(cb))
}
