package middlewares

import (
	"overengineering-my-application/app/infrastructure/resilience/observers/ratelimiter"
	"overengineering-my-application/app/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleware(config *util.Config) middleware.KeyAuthValidator {
	rl := ratelimiter.NewRateLimiter(config.RateLimiter, config.RateLimiterExponentialBaseFactor)

	return func(key string, c echo.Context) (bool, error) {
		err := rl.Call(func() error {
			return nil
		})

		return err == nil, err
	}
}
