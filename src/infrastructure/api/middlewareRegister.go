package api

import (
	"observability-series-golang-edition/src/infrastructure/metrics"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRegister(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(metrics.PrometheusMiddleware)
}
