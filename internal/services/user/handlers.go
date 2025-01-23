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
		Account:  payload.Account,
		Email:    payload.Email,
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

}
