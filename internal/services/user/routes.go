package user

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store       store.UserRepository
	middlewares *middlewares.Middleware
}

func NewHandler(store store.UserRepository, middlewares *middlewares.Middleware) *Handler {
	return &Handler{store, middlewares}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Post("/register", h.postRegister)
	r.Post("/login", h.postLogin)
}
