package product

import (
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (h *Handler) postProducts(w http.ResponseWriter, r *http.Request) {
	payload, err := h.store.ParseAddProductPayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.ValidateAddProductPayload(payload); err != nil {
		myhttp.SendError(w, err)
		return
	}

	idString, err := h.globalStore.GetSingleClaimFromContext(r, "sub")
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	id, err := h.globalStore.Atoi(idString)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	err = h.store.AddProduct(&models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       *payload.Price,
		Stock:       *payload.Stock,
		CreatedBy:   id,
		UpdateBy:    id,
	})
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusCreated, "product created successfully")
}
