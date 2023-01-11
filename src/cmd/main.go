package main

import (
	"encoding/json"
	"fmt"
	"golang-cicd-workflow-heroku/src/util"
	"net/http"
)

func loadEnv() util.Config {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic("Error loading config: " + err.Error())
	}

	return config
}

func main() {
	config := loadEnv()

	fmt.Println("Server is running on port " + config.Port)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("ok")
	})

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
