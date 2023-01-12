package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheckHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, "ok")
}
