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
	// login
	r.Post("/register", h.postRegister)
	r.Post("/login", h.postLogin)

	r.With(h.middls.AuthWithClaims).Post("/logout", h.postLogout)

	// profile
	r.With(h.middls.AuthWithClaims).Get("/profile", h.getProfile)

	r.With(h.middls.AuthWithClaims).Patch("/profile", h.patchProfile)

	// users.. <-- admin urls
	r.With(h.middls.AuthAdmin).Get("/users", h.getUsers)

	r.With(h.middls.AuthAdmin).Get("/users/{id}", h.getUserById)
	r.With(h.middls.AuthAdmin).Get("/users/role/{role}", h.getUsersByRole)

	r.With(h.middls.AuthAdmin).Patch("/users/{id}/role", h.patchRoleUser)
}
