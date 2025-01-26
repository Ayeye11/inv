package inventory

import (
	"github.com/Ayeye11/inv/internal/middlewares"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	GlobalRepository store.GlobalRepository
	store            store.InventoryRepository
	middls           *middlewares.Middleware
}

func NewHandler(globalStore store.GlobalRepository, store store.InventoryRepository, middls *middlewares.Middleware) *Handler {
	return &Handler{globalStore, store, middls}
}

func (h *Handler) SetRoutes(r *chi.Mux) {
	r.With(h.middls.AuthEmployee).Get("/inventory", h.showInventory)
}
