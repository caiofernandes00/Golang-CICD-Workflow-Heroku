package middlewares

import (
	"errors"
	"overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CircuitBreakerMiddleware(cb *circuitbreaker.CircuitBreaker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := cb.Call(func() error {
				if err := next(c); err != nil {
					return err
				}
				statusCode := strconv.Itoa(c.Response().Status)
				if statusCode == "500" {
					return errors.New("error")
				}

				return nil
			})

			return err
		}
	}
}
