package api

import (
	"golang-cicd-workflow-heroku/src/infrastructure/metrics"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(metrics.PrometheusMiddleware)
}
