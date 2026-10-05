package main

import (
  "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

)

type TestResponse struct {
	Status string `json:"status"`
}

func main() {
  r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request){
		body := TestResponse{Status: "ok"}
		render.JSON(w, r, body)
	})
	http.ListenAndServe(":3000", r)
}