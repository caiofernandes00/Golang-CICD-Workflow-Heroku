package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"overengineering-my-application/app/infrastructure/api"
	"overengineering-my-application/app/infrastructure/circuitbreaker"
	"overengineering-my-application/app/infrastructure/metrics"
	"overengineering-my-application/app/util"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"

	_ "overengineering-my-application/docs"
)

var (
	e           *echo.Echo
	http2Server *http2.Server
	config      *util.Config
)

func init() {
	e = echo.New()
	loadEnv()
	cb := circuitbreaker.NewCircuitBreaker(config.CircuitBreakerInterval, config.CircuitBreakerThreshold)
	metrics.MetricsRegister()
	api.MiddlewareRegister(e, config, cb, loggerSetup(), gzipSetup(config))
	api.RoutesRegister(e)
	loadHttp2Server()
}

// @title Overengineering My Application API
// @version 1.0
// @description Example Golang REST API
// @contact.name Caio Fernandes
// @contact.url https://github.com/caiofernandes00
// @contact.email caiow.wk@gmail.com
// @BasePath /api/v1
func main() {
	go func() {
		if err := e.StartH2CServer(":"+config.Port, http2Server); err != http.ErrServerClosed {
			log.Fatal(err)
		}

		log.Println("Server is running on port " + config.Port)
	}()

	gracefulShutdown()
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func loadHttp2Server() {
	http2Server = &http2.Server{
		MaxConcurrentStreams: config.MaxConcurrentStreams,
		MaxReadFrameSize:     config.MaxReadFrameSize,
		IdleTimeout:          config.IdleTimeout * time.Second,
	}
}

func loadEnv() {
	config = util.NewConfig()
	path, err := getRootFile()

	if err == nil {
		_ = config.LoadEnvFile(path)
		return
	}

	_, _ = config.LoadEnv()
}

func getRootFile() (ex string, err error) {
	ex, _ = os.Getwd()
	_, err = os.Stat(filepath.Join(ex, "app.env"))

	if err != nil {
		ex = filepath.Join(ex, "../../")
		_, err = os.Stat(filepath.Join(ex, "app.env"))

		if err != nil {
			log.Println("No env file provided, using only env variables")
		}
	}

	return
}

func loggerSetup() middleware.RequestLoggerConfig {
	logger := zerolog.New(os.Stdout)

	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}
}

func gzipSetup(config *util.Config) middleware.GzipConfig {
	return middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			for _, s := range config.SkipCompressionUrls {
				if strings.Contains(c.Request().URL.Path, s) {
					return true
				}
			}
			return false
		},
	}
}
