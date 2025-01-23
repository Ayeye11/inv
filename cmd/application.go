package main

import (
	"log"
	"net/http"

	"github.com/Ayeye11/inv/internal/store"
)

type application struct {
	addr    string
	storage store.Storage
}

func newApplication(addr string, storage store.Storage) *application {
	return &application{addr, storage}
}

func (app *application) run() error {
	server := http.Server{
		Addr:    app.addr,
		Handler: nil,
	}

	log.Printf("server run on %s", server.Addr)
	return server.ListenAndServe()
}
