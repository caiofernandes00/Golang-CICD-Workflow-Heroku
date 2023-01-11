package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Println("Server is running on port " + port)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		json.NewDecoder(r.Body).Decode(&req)
		res := Response{Message: "Hello " + req.Name}
		json.NewEncoder(w).Encode(res)
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
