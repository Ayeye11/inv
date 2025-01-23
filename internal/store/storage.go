package store

import (
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/golang-jwt/jwt/v5"
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
	// auth
	GetClaimsFromCookie(r *http.Request) (jwt.MapClaims, error)
	GetSingleClaim(claims jwt.MapClaims, key string) (string, error)
	CheckRole(role, minRole string) error
}

type UserRepository interface {
	// parse
	ParseRegisterPayload(r *http.Request) (*models.UserRegisterPayload, error)
	ParseLoginPayload(r *http.Request) (*models.UserLoginPayload, error)
	// validate
	ValidateRegisterPayload(payload *models.UserRegisterPayload) error
	ValidateLoginPayload(payload *models.UserLoginPayload) error
	// auth
	HashUserPassword(password string) (string, error)
	TryLogin(email, password string) (*models.User, error)
	CreateToken(user *models.User) (string, error)
	SendCookie(w http.ResponseWriter, token string)
	CleanCookie(w http.ResponseWriter)
	// create
	CreateUser(user *models.User) error
	// read
	GetUserByEmail(email string) (*models.User, error)
}
