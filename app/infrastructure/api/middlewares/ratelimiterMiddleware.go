package middlewares

import (
	"overengineering-my-application/app/infrastructure/resilience/observers/ratelimiter"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleware(rt time.Duration) middleware.KeyAuthValidator {
	rl := ratelimiter.NewRateLimiter(rt)

	return func(key string, c echo.Context) (bool, error) {
		err := rl.Call(func() error {
			return nil
		})

		return err == nil, err
	}
}
