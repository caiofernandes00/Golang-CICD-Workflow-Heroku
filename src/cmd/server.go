package main

import (
	"crypto/tls"
	"fmt"
	"golang-cicd-workflow-heroku/src/infrastructure/api"
	"golang-cicd-workflow-heroku/src/infrastructure/metrics"
	"golang-cicd-workflow-heroku/src/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
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

	autoTLSManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("/var/www/.cache"),
		HostPolicy: autocert.HostWhitelist(config.HostPolicy),
	}
	s := http.Server{
		Addr:    ":" + config.Port,
		Handler: e,
		TLSConfig: &tls.Config{
			GetCertificate: autoTLSManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
		ReadTimeout: 30 * time.Second,
	}

	if err := s.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
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
