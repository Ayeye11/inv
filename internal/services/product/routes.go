package product

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	globalStore store.GlobalRepository
	store       store.ProductRepository
	middl       *middlewares.Middleware
}

func NewHandler(globalStore store.GlobalRepository, store store.ProductRepository, middl *middlewares.Middleware) *Handler {
	return &Handler{globalStore, store, middl}
}

func (h *Handler) SetRoutes(r *chi.Mux) {
	r.With(h.middl.AuthorizeAdmin).Post("/products", h.postProducts)
}
