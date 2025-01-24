package router

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/services/product"
	"github.com/Ayeye11/inv/internal/services/user"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	mux     *chi.Mux
	storage store.Storage
}

func NewRouter(mux *chi.Mux, storage store.Storage) *Router {
	return &Router{mux, storage}
}

func (r *Router) Setup() {
	middls := middlewares.SetMiddlewares(r.storage.Middleware)

	// handlers
	authHandler := user.NewHandler(r.storage.Global, r.storage.User, middls)
	productHandler := product.NewHandler(r.storage.Global, r.storage.Product, middls)

	// routes
	authHandler.SetupRoutes(r.mux)
	productHandler.SetRoutes(r.mux)
}
