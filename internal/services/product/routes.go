package product

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	globalStore store.GlobalRepository
	store       store.ProductRepository
	middls      *middlewares.Middleware
}

func NewHandler(globalStore store.GlobalRepository, store store.ProductRepository, middls *middlewares.Middleware) *Handler {
	return &Handler{globalStore, store, middls}
}

func (h *Handler) SetRoutes(r *chi.Mux) {
	r.With(h.middls.AuthEmployeeWithClaims).Post("/products", h.postProducts)

	r.With(h.middls.AuthEmployee).Get("/products", h.getProductsPage)
	r.With(h.middls.AuthEmployee).Get("/products/{id}", h.getProductById)
}
