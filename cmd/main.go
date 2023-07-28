package main

import (
	"log"
	"net/http"

	"github.com/mtolley/hello-api/handlers"
	"github.com/mtolley/hello-api/handlers/rest"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
