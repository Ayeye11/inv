package user

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	globalStore store.GlobalRepository
	store       store.UserRepository
	middls      *middlewares.Middleware
}

func NewHandler(globalStore store.GlobalRepository, store store.UserRepository, middls *middlewares.Middleware) *Handler {
	return &Handler{globalStore, store, middls}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Post("/register", h.postRegister)
	r.Post("/login", h.postLogin)

	r.With(h.middls.AuthWithClaims).Post("/logout", h.postLogout)
}
