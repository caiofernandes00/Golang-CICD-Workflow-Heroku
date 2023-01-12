package api

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", HealthCheckHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
