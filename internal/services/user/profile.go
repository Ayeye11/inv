package user

import (
	"net/http"

	"github.com/Ayeye11/inv/pkg/myhttp"
)

// get
func (h *Handler) getProfile(w http.ResponseWriter, r *http.Request) {
	val, err := h.globalStore.GetSingleClaimFromContext(r, "sub")
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	id, err := h.globalStore.Atoi(val)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	res := user.ToShowProfile()

	myhttp.SendJSON(w, http.StatusOK, res)
}

func (h *Handler) patchProfile(w http.ResponseWriter, r *http.Request) {
	val, err := h.globalStore.GetSingleClaimFromContext(r, "sub")
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	id, err := h.globalStore.Atoi(val)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	payload, err := h.store.ParseUserUpdatePayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	updates, err := h.store.ValidatePatchUser(payload)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.PatchUser(user.ID, updates); err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusOK, "user updated successfully")
}
