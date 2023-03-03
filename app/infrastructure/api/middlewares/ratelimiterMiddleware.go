package middlewares

import (
	"overengineering-my-application/app/infrastructure/resilience"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleware(rateLimiter time.Duration) middleware.KeyAuthValidator {
	rl := resilience.NewRateLimiter(rateLimiter)

	return func(key string, c echo.Context) (bool, error) {
		allowed, err := rl.Allow()
		if err != nil {
			return false, resilience.ErrRateLimitExceeded
		}

		return allowed, nil
	}
}
