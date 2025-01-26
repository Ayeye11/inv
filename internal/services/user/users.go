package user

import (
	"fmt"
	"net/http"

	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("page") {
		http.Redirect(w, r, "/users?page=1", http.StatusFound)
		return
	}

	query := r.URL.Query().Get("page")
	page, err := h.globalStore.Atoi(query)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	users, err := h.store.GetUsersByRolePage("all", page)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, users)
}

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("id")
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

	myhttp.SendJSON(w, http.StatusOK, user.ToShowProfile())
}

func (h *Handler) getUsersByRole(w http.ResponseWriter, r *http.Request) {
	role := r.PathValue("role")
	if err := h.store.CheckRole(role); err != nil {
		myhttp.SendError(w, err)
		return
	}

	if !r.URL.Query().Has("page") {
		http.Redirect(w, r, fmt.Sprintf("/users/role/%s?page=1", role), http.StatusFound)
		return
	}

	query := r.URL.Query().Get("page")
	page, err := h.globalStore.Atoi(query)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	users, err := h.store.GetUsersByRolePage(role, page)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, users)
}

// update
// to fix.. illegible
func (h *Handler) patchRoleUser(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("id")
	id, err := h.globalStore.Atoi(val)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	payload, err := h.store.ParseUpdateRole(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if user.Role == "admin" {
		myhttp.SendError(w, myhttp.NewErrorHTTP(http.StatusForbidden, "you cannot change the role of this user"))
		return
	}

	if err := h.store.PatchRoleUser(id, payload.Role); err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusOK, "user updated successfully")
}
