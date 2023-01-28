package main

import (
	"fmt"
	"log"
	"observability-series-golang-edition/src/infrastructure/api"
	"observability-series-golang-edition/src/infrastructure/metrics"
	"observability-series-golang-edition/src/util"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var (
	e      *echo.Echo
	config util.Config
)

func init() {
	e = echo.New()
	config = loadEnv()
	metrics.MetricsRegister()
	api.MiddlewareRegister(e)
	api.RoutesRegister(e)
}

func main() {
	e.Logger.Fatal(e.Start(":" + config.Port))
}

func loadEnv() util.Config {
	config, _ := util.LoadConfig(getRootFile())

	fmt.Println("Server is running on port " + config.Port)
	return config
}

func getRootFile() string {
	ex, _ := os.Getwd()
	_, err := os.Stat(filepath.Join(ex, "app.env"))
	if err != nil {
		ex = filepath.Join(ex, "../../")
		_, err = os.Stat(filepath.Join(ex, "app.env"))
		if err != nil {
			log.Fatal("Error loading config: " + err.Error())
			panic(err)
		}
	}

	return ex
}
