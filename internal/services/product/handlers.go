package product

import (
	"fmt"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
)

// post
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
		Category:    payload.Category,
		CreatedBy:   id,
		UpdatedBy:   id,
	})
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusCreated, "product created successfully")
}

// get
func (h *Handler) getProductsPage(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("page") {
		http.Redirect(w, r, "/products?page=1", http.StatusFound)
		return
	}

	query := r.URL.Query().Get("page")

	page, err := h.globalStore.Atoi(query)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	products, err := h.store.GetProductsPage(page)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, products)
}

func (h *Handler) getProductById(w http.ResponseWriter, r *http.Request) {
	id, err := h.globalStore.Atoi(r.PathValue("id"))
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	product, err := h.store.GetProductById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, product)
}

func (h *Handler) getProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")

	if !r.URL.Query().Has("page") {
		url := fmt.Sprintf("/products/%s?page=1", category)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	query := r.URL.Query().Get("page")
	page, err := h.globalStore.Atoi(query)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	prods, err := h.store.GetProductsByCategoryPage(page, category)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendJSON(w, http.StatusOK, prods)
}

// update
func (h *Handler) putProductById(w http.ResponseWriter, r *http.Request) {
	id, err := h.globalStore.Atoi(r.PathValue("id"))
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	payload, err := h.store.ParseUpdateProductPayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.ValidatePutUpdate(payload); err != nil {
		myhttp.SendError(w, err)
		return
	}

	prod, err := h.store.GetProductById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	claim, err := h.globalStore.GetSingleClaimFromContext(r, "sub")
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	userID, err := h.globalStore.Atoi(claim)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	err = h.store.UpdatePutProduct(prod.ID, &models.Product{
		Name:        *payload.Name,
		Description: payload.Description,
		Price:       *payload.Price,
		Stock:       *payload.Stock,
		Category:    *payload.Category,
		UpdatedBy:   userID,
	})
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusOK, "product updated successfully")
}

func (h *Handler) patchProductById(w http.ResponseWriter, r *http.Request) {
	id, err := h.globalStore.Atoi(r.PathValue("id"))
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	payload, err := h.store.ParseUpdateProductPayload(r)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	updates, err := h.store.ValidatePatchUpdate(payload)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	prod, err := h.store.GetProductById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	claim, err := h.globalStore.GetSingleClaimFromContext(r, "sub")
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	userID, err := h.globalStore.Atoi(claim)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.UpdatePatchProduct(prod.ID, updates, userID); err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusOK, "product updated successfully")
}

// delete
func (h *Handler) deleteProductById(w http.ResponseWriter, r *http.Request) {
	id, err := h.globalStore.Atoi(r.PathValue("id"))
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	_, err = h.store.GetProductById(id)
	if err != nil {
		myhttp.SendError(w, err)
		return
	}

	if err := h.store.DeleteProductById(id); err != nil {
		myhttp.SendError(w, err)
		return
	}

	myhttp.SendMessage(w, http.StatusOK, "product deleted successfully")
}
