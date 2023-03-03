package middleware

import (
	"bytes"
	"log"
	"net/http"
	"overengineering-my-application/app/infrastructure/cache"
	"overengineering-my-application/app/util"

	"github.com/labstack/echo/v4"
)

func CacheMiddleware(config *util.Config) echo.MiddlewareFunc {
	var cache = cache.NewLRUCache[any](config.CacheRequestCapacity)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cacheKey := c.Request().URL.String()
			if v, ok := cache.Get(cacheKey); ok {
				log.Println("Cache hit: " + cacheKey)
				_, _ = c.Response().Write([]byte(v.(string)))
				return nil
			}

			buf := new(bytes.Buffer)
			rw := c.Response().Writer
			c.Response().Writer = &responseWriter{rw, buf}

			if err := next(c); err != nil {
				c.Error(err)
			}

			cacheHeader := c.Response().Header().Get("Cache-Control")
			if cacheHeader == "" || cacheHeader == "no-cache" || cacheHeader == "no-store" {
				return nil
			} else {
				log.Println("Cache set: " + cacheKey)
				cache.Set(cacheKey, buf.String(), config.CacheRequestTTL)
			}
			return nil
		}
	}
}

type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	written, err := rw.body.Write(b)
	if err != nil {
		return written, err
	}

	return rw.ResponseWriter.Write(b)
}
