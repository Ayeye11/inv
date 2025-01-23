package migrate

import (
	"github.com/Ayeye11/inv/internal/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
