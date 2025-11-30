package main

import (
	"net/http"

	"github.com/rudnero/go-musthave-metrica.git/internal/handler"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/update/{type}/{name}/{value}", handler.UpdateHandle)
	return http.ListenAndServe(`:8080`, mux)
}
