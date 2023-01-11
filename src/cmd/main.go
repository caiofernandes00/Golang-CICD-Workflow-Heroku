package main

import (
	"encoding/json"
	"fmt"
	"golang-cicd-workflow-heroku/src/util"
	"log"
	"net/http"
)

func loadEnv() util.Config {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: " + err.Error())
	}

	return config
}

func main() {
	config := loadEnv()
	// db.Connect(config)

	fmt.Println("Server is running on port " + config.Port)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("ok")
	})

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal("Error starting server: " + err.Error())
	}
}
