package api

import (
	"net/http"
	"time"

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

func UsersRouteHandler(e echo.Context) error {
	time.Sleep(5 * time.Second)
	return e.JSON(http.StatusOK, map[string]string{"name": "John", "surname": "Doe"})
}
