package store

import (
	"errors"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

// parse
func (s *ProductStore) ParseAddProductPayload(r *http.Request) (*models.AddProductPayload, error) {
	var payload models.AddProductPayload

	if err := myhttp.ParseJSON(r, &payload); err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid request")
	}

	return &payload, nil
}

// validate
func (s *ProductStore) ValidateAddProductPayload(payload *models.AddProductPayload) error {
	if payload == nil {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing request")
	}

	if payload.Name == "" || payload.Price == nil || payload.Stock == nil {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing credentials")
	}

	if len(payload.Name) > 100 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	if payload.Description != nil && len(*payload.Description) > 255 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	return nil
}

// create
func (s *ProductStore) AddProduct(product *models.Product) error {
	if err := s.db.Create(&product).Error; err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return myhttp.NewErrorHTTP(http.StatusConflict, "product already exists")
		}

		return myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// read
func (s *ProductStore) GetProductsPage(page int) ([]models.Product, error) {
	if page < 1 {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid query")
	}

	limit := 10
	offset := (page - 1) * limit

	query := s.db.Limit(limit).Offset(offset)

	var products []models.Product
	if err := query.Find(&products).Error; err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	if page > 1 && len(products) == 0 {
		return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "no products found")
	}

	return products, nil
}

func (s *ProductStore) GetProductById(id int) (*models.Product, error) {
	if id < 1 {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid id")
	}

	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "product not found")
		}

		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return &product, nil
}
