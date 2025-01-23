package user

import (
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (h *Handler) postRegister(w http.ResponseWriter, r *http.Request) {
	payload, err := h.store.ParseRegisterPayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.ValidateRegisterPayload(payload); err != nil {
		myhttp.SendError(w, err)
		return
	}

	hash, err := h.store.HashUserPassword(payload.Password)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	err = h.store.CreateUser(&models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Lastname: payload.Lastname,
		Password: hash,
		Role:     "user",
	})
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusCreated, "account created successfully")
}

func (h *Handler) postLogin(w http.ResponseWriter, r *http.Request) {
	payload, err := h.store.ParseLoginPayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.ValidateLoginPayload(payload); err != nil {
		myhttp.SendError(w, err)
		return
	}

	user, err := h.store.TryLogin(payload.Email, payload.Password)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	token, err := h.store.CreateToken(user)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	h.store.SendCookie(w, token)
	myhttp.SendMessage(w, http.StatusOK, "login successful")
}
