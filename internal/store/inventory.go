package store

import (
	"math"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
	"gorm.io/gorm"
)

type InventoryStore struct {
	db *gorm.DB
}

func (s *InventoryStore) InventoryValue() (*models.Inventory, error) {
	var data models.Inventory
	var products []models.Product
	if err := s.db.Select("stock", "price").Find(&products).Error; err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	var total float64

	for _, item := range products {
		data.Products++
		total = float64(item.Stock) * item.Price
		data.Value += math.Round(total*100) / 100
	}

	return &data, nil
}
