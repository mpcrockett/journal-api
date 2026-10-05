package main

import (
	"encoding/json"
	"log"
  "net/http"
)

type TestResponse struct {
	Status string `json:"status"`
}

func main() {
  mux := http.NewServeMux()

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		res := TestResponse{
			Status: "ok",
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	mux.HandleFunc("/health", h1)

	log.Print("listening...")

	http.ListenAndServe(":8080", mux)
}