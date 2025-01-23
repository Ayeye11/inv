package store

import (
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"gorm.io/gorm"
)

type Storage struct {
	User       UserRepository
	Middleware MiddlewareRepository
}

func NewStorage(db *gorm.DB, jwtKey string) Storage {
	return Storage{
		&UserStore{db, jwtKey},
		&MiddlewareStore{jwtKey},
	}
}

type MiddlewareRepository interface {
}

type UserRepository interface {
	// parse
	ParseRegisterPayload(r *http.Request) (*models.UserRegisterPayload, error)
	// validate
	ValidateRegisterPayload(payload *models.UserRegisterPayload) error
	// auth
	HashUserPassword(password string) (string, error)
	CreateToken(user *models.User) (string, error)
	// create
	CreateUser(user *models.User) error
}
