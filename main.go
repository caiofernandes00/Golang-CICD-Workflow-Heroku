package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on port " + port)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("ok")
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
