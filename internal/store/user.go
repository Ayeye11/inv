package store

import (
	"errors"
	"net/http"
	"time"

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
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid request")
	}

	return &payload, nil
}

func (s *UserStore) ParseLoginPayload(r *http.Request) (*models.UserLoginPayload, error) {
	var payload models.UserLoginPayload
	if err := myhttp.ParseJSON(r, &payload); err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid request")
	}

	return &payload, nil
}

// validate
func (s *UserStore) ValidateRegisterPayload(payload *models.UserRegisterPayload) error {
	if payload.Email == "" || payload.Name == "" || payload.Lastname == "" || payload.Password == "" {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing credentials")
	}

	if !myhttp.ValidateEmail(payload.Email) || len(payload.Password) < 7 || len(payload.Password) > 100 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	}

	if len(payload.Name) > 100 || len(payload.Lastname) > 100 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid name")
	}

	return nil
}

func (*UserStore) ValidateLoginPayload(payload *models.UserLoginPayload) error {
	if payload.Email == "" || payload.Password == "" {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing credentials")
	}

	if !myhttp.ValidateEmail(payload.Email) || len(payload.Password) < 7 || len(payload.Password) > 100 {
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

func (s *UserStore) TryLogin(email, password string) (*models.User, error) {
	user, err := s.GetUserByEmail(email)
	if err != nil {

		errHTTP := myhttp.ParseError(err)
		if errHTTP.Status == 404 {
			return nil, myhttp.NewErrorHTTP(http.StatusUnauthorized, "invalid email or password")
		}

		return nil, errHTTP
	}

	if !auth.ComparePassword(user.Password, password) {
		return nil, myhttp.NewErrorHTTP(http.StatusUnauthorized, "invalid email or password")
	}

	return user, nil
}

func (s *UserStore) CreateToken(user *models.User) (string, error) {
	token, err := auth.CreateJWT(user, s.jwtKey)
	if err != nil {
		return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail to create jwt")
	}

	return token, nil
}

func (s *UserStore) SendCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(time.Minute * 30),
	})
}

func (s *UserStore) CleanCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
	})
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

// read
func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "user doesn't exists")
		}

		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}
