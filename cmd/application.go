package main

import (
	"log"
	"net/http"

	"github.com/Ayeye11/inv/internal/router"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type application struct {
	addr    string
	storage store.Storage
}

func newApplication(addr string, storage store.Storage) *application {
	return &application{addr, storage}
}

func (app *application) run() error {
	r := chi.NewRouter()

	router := router.NewRouter(r, app.storage)
	router.Setup()

	server := http.Server{
		Addr:    app.addr,
		Handler: r,
	}

	log.Printf("server run on %s", server.Addr)
	return server.ListenAndServe()
}
