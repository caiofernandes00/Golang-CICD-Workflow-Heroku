package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// @Summary Health check
// @Description Health check for the application
// @Tags infra
// @Accept  json
// @Produce  json
// @Success 200 {string}
// @Router /users [get]
func HealthCheckRouteHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, "ok")
}

// @Summary Generate bad request error
// @Description Generate bad request error - 400 Bad Request
// @Tags errors
// @Accept  json
// @Produce  json
// @Error 400 {object} map[string]string
// @Router /bad_request_error [get]
func BadRequestErrorRouteHandler(e echo.Context) error {
	return e.JSON(http.StatusBadRequest, map[string]string{"error": "bad request error"})
}

// @Summary Generate handled internal error
// @Description Generate handled internal error
// @Tags errors
// @Accept  json
// @Produce  json
// @Error 500 {object} map[string]string
// @Router /internal_error [get]
func InternalErrorRouteHandler(e echo.Context) error {
	return e.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
}

// @Summary Generate unhandled error
// @Description Generate unhandled error
// @Tags errors
// @Accept  json
// @Produce  json
// @Error 500 {object} map[string]string
// @Router /unexpected_error [get]
func UnexpectedErrorRouteHandler(e echo.Context) error {
	return errors.New("unexpected error")
}

// @Summary Get users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /users [get]
func UsersRouteHandler(e echo.Context) error {
	time.Sleep(5 * time.Second)
	e.Response().Header().Set("Cache-Control", "public, max-age=60")
	return e.JSON(http.StatusOK, map[string]string{"name": "John", "surname": "Doe"})
}
