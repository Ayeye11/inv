package inventory

import (
	"net/http"

	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (h *Handler) showInventory(w http.ResponseWriter, r *http.Request) {
	data, err := h.store.InventoryValue()
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, data)
}
