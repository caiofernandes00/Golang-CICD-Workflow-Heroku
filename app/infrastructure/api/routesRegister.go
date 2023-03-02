package api

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RoutesRegister(e *echo.Echo) {
	e.GET("/health", HealthCheckRouteHandler)
	e.GET("/bad_request_error", BadRequestErrorRouteHandler)
	e.GET("/internal_error", InternalErrorRouteHandler)
	e.GET("/unexpected_error", UnexpectedErrorRouteHandler)
	// @Summary Get users
	// @Description Get all users
	// @Tags users
	// @ID get-users
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} UserResponse
	// @Router /users [get]
	e.GET("/users", UsersRouteHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
