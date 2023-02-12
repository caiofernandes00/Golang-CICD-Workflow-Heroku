package main

import (
	"log"
	"net/http"
	"observability-series-golang-edition/app/infrastructure/api"
	"observability-series-golang-edition/app/infrastructure/metrics"
	"observability-series-golang-edition/app/util"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

var (
	e           *echo.Echo
	http2Server *http2.Server
	config      util.Config
)

func init() {
	e = echo.New()
	config = loadEnv()
	metrics.MetricsRegister()
	api.MiddlewareRegister(e)
	api.RoutesRegister(e)
	loadHttp2Server()
}

func main() {
	if err := e.StartH2CServer(":"+config.Port, http2Server); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	log.Println("Server is running on port " + config.Port)
}

func loadHttp2Server() {
	http2Server = &http2.Server{
		MaxConcurrentStreams: config.MaxConcurrentStreams,
		MaxReadFrameSize:     config.MaxReadFrameSize,
		IdleTimeout:          config.IdleTimeout * time.Second,
	}
}

func loadEnv() (config util.Config) {
	path, err := getRootFile()

	if err == nil {
		config, _ = util.LoadEnvFile(path)
	}

	util.LoadEnv()

	return
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
