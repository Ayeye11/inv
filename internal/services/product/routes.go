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

// prefix "/products"
func (h *Handler) SetRoutes(r *chi.Mux) {
	r.With(h.middls.AuthEmployeeWithClaims).Post("/", h.postProducts)

	r.With(h.middls.Auth).Get("/", h.getProductsPage) // <-- /?page=x
	r.With(h.middls.Auth).Get("/{id}", h.getProductById)
	r.With(h.middls.Auth).Get("/{category}", h.getProductsByCategory)

	r.With(h.middls.AuthEmployeeWithClaims).Put("/{id}", h.putProductById)
	r.With(h.middls.AuthEmployeeWithClaims).Patch("/{id}", h.patchProductById)

	r.With(h.middls.AuthEmployee).Delete("/{id}", h.deleteProductById)
}
