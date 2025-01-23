package router

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/services/user"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	r       *chi.Mux
	storage store.Storage
}

func NewRouter(r *chi.Mux, storage store.Storage) *Router {
	return &Router{r, storage}
}

func (mux *Router) Setup() {
	middls := middlewares.SetMiddlewares(mux.storage.Middleware)

	// handlers
	authHandler := user.NewHandler(mux.storage.User, middls)

	// routes
	authHandler.SetupRoutes(mux.r)
}
