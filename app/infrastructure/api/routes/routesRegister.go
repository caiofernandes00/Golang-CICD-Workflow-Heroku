package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RoutesRegister(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/health", HealthCheckRouteHandler)
	v1.GET("/bad_request_error", BadRequestErrorRouteHandler)
	v1.GET("/internal_error", InternalErrorRouteHandler)
	v1.GET("/unexpected_error", UnexpectedErrorRouteHandler)
	v1.GET("/users", UsersRouteHandler)

	v1.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	v1.GET("/swagger/*", echoSwagger.WrapHandler)
}
