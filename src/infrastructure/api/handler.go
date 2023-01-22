package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheckRouteHandler(e echo.Context) error {
	return e.String(http.StatusOK, "ok")
}

func BadRequestErrorRouteHandler(e echo.Context) error {
	return e.String(http.StatusBadRequest, "error")
}

func InternalErrorRouteHandler(e echo.Context) error {
	return e.String(http.StatusInternalServerError, "error")
}
