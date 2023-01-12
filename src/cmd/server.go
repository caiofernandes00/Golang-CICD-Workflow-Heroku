package main

import (
	"fmt"
	"golang-cicd-workflow-heroku/src/infrastructure/api"
	"golang-cicd-workflow-heroku/src/infrastructure/metrics"
	"golang-cicd-workflow-heroku/src/util"
	"log"

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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: " + err.Error())
	}

	fmt.Println("Server is running on port " + config.Port)
	return config
}
