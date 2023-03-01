package cache

import (
	"bytes"
	"net/http"

	"github.com/labstack/echo/v4"
)

var cache = NewLRUCache[any](100, 100)

func CacheMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cacheKey := c.Request().URL.String()
		if v, ok := cache.Get(cacheKey); ok {
			c.Response().Write([]byte(v.(string)))
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

			cache.Set(cacheKey, buf.String())
		}
		return nil
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
