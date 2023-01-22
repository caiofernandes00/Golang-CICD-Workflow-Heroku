package main

import (
	"fmt"
	"golang-cicd-workflow-heroku/src/infrastructure/api"
	"golang-cicd-workflow-heroku/src/infrastructure/metrics"
	"golang-cicd-workflow-heroku/src/util"
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo
var config util.Config

func init() {
	e = echo.New()
	config = loadEnv()
	metrics.RegisterMetrics()
	api.RegisterMiddleware(e)
	api.RegisterRoutes(e)
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
