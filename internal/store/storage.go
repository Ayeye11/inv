package store

import (
	"context"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Storage struct {
	Global     GlobalRepository
	Middleware MiddlewareRepository
	User       UserRepository
	Product    ProductRepository
}

func NewStorage(db *gorm.DB, jwtKey string) Storage {
	return Storage{
		Global:     &GlobalStore{},
		Middleware: &MiddlewareStore{jwtKey},
		User:       &UserStore{db, jwtKey},
		Product:    &ProductStore{db},
	}
}

type GlobalRepository interface {
	// parse
	Atoi(x string) (int, error)
	// context
	GetClaimsFromContext(r *http.Request) (jwt.MapClaims, error)
	GetSingleClaimFromContext(r *http.Request, key string) (string, error)
}

type MiddlewareRepository interface {
	// auth
	GetClaimsFromCookie(r *http.Request) (jwt.MapClaims, error)
	GetSingleClaim(claims jwt.MapClaims, key string) (string, error)
	CheckRole(role, minRole string) error
	// context
	SetClaimsToContext(r *http.Request, claims jwt.MapClaims) context.Context
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
	// cookie
	SendCookie(w http.ResponseWriter, token string)
	CleanCookie(w http.ResponseWriter)
	// create
	CreateUser(user *models.User) error
	// read
	GetUserByEmail(email string) (*models.User, error)
}

type ProductRepository interface {
	// parse
	ParseAddProductPayload(r *http.Request) (*models.AddProductPayload, error)
	// validate
	ValidateAddProductPayload(payload *models.AddProductPayload) error
	// create
	AddProduct(product *models.Product) error
}
