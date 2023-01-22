package api

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", HealthCheckRouteHandler)
	e.GET("/bad_request_error", BadRequestErrorRouteHandler)
	e.GET("/internal_error", InternalErrorRouteHandler)
	e.GET("/users", UsersRouteHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
