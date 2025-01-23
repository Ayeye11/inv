package store

import (
	"errors"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/internal/utils/auth"
	"github.com/Ayeye11/inv/pkg/myhttp"
	"gorm.io/gorm"
)

type UserStore struct {
	db     *gorm.DB
	jwtKey string
}

// parse
func (s *UserStore) ParseRegisterPayload(r *http.Request) (*models.UserRegisterPayload, error) {
	var payload models.UserRegisterPayload
	if err := myhttp.ParseJSON(r, &payload); err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	return &payload, nil
}

// validate
func (s *UserStore) ValidateRegisterPayload(payload *models.UserRegisterPayload) error {
	if payload.Email == "" || payload.Account == "" || payload.Password == "" {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing credentials")
	}

	if !myhttp.ValidateEmail(payload.Email) || len(payload.Account) < 5 || len(payload.Password) < 7 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	if len(payload.Account) > 50 || len(payload.Password) > 100 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	return nil
}

// auth
func (s *UserStore) HashUserPassword(password string) (string, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail hash password")
	}

	return hash, nil
}

func (s *UserStore) CreateToken(user *models.User) (string, error) {
	token, err := auth.CreateJWT(user, s.jwtKey)
	if err != nil {
		return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail to create jwt")
	}

	return token, nil
}

// create
func (s *UserStore) CreateUser(user *models.User) error {
	if err := s.db.Create(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return myhttp.NewErrorHTTP(http.StatusConflict, "account already exists")

		}
		return myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return nil
}
