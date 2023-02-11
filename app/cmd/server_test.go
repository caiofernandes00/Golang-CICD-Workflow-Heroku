package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Endpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `ok`)
	}))

	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("Expected content type %s, got %s", "text/plain", resp.Header.Get("Content-Type"))
	}
}
