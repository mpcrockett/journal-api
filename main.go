package main

import (
	"log"
  "net/http"
)

type TestResponse struct {
	Status string `json:"status"`
}

func main() {
  mux := http.NewServeMux()
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		log.Print("Hello World!")
	}
	mux.HandleFunc("/health", h1)
	log.Print("listening...")
	http.ListenAndServe(":8080", mux)
}