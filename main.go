package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := Response{Message: "Hello " + req.Name}
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe(":8080", nil)
}
