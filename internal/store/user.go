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

func (s *UserStore) ParseUserUpdatePayload(r *http.Request) (*models.UserUpdatePayload, error) {
	var payload models.UserUpdatePayload
	if err := myhttp.ParseJSON(r, &payload); err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid request")
	}

	return &payload, nil
}

// validate
func (s *UserStore) ValidateRegisterPayload(payload *models.UserRegisterPayload) error {
	if payload == nil {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "missing request")
	}

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

func (s *UserStore) CheckRole(role string) error {
	validRoles := map[string]bool{
		"admin":    true,
		"employee": true,
		"user":     true,
	}

	if exists := validRoles[role]; !exists {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid role")
	}

	return nil
}

func (*UserStore) ValidatePatchUser(payload *models.UserUpdatePayload) (map[string]any, error) {
	if payload == nil {
		myhttp.NewErrorHTTP(http.StatusBadRequest, "missing request")
	}

	var response = make(map[string]any)

	if payload.Email != nil {
		response["email"] = *payload.Email
	}

	if payload.Name != nil {
		response["name"] = *payload.Name
	}

	if payload.Lastname != nil {
		response["lasname"] = *payload.Lastname
	}

	if payload.Password != nil {
		response["password"] = *payload.Password
	}

	return response, nil
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

// cookie
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
			return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "user not found")
		}

		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}

func (s *UserStore) GetUserById(id int) (*models.User, error) {
	if id < 1 {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid id")
	}

	var user models.User

	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "user not found")
		}

		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}

// update
func (s *UserStore) PatchUser(id int, updates map[string]any) error {
	if id < 1 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid id")
	}

	if password, exists := updates["password"].(string); exists {
		hash, err := s.HashUserPassword(password)
		if err != nil {
			return err
		}

		updates["password"] = hash
	}

	if err := s.db.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// --- ADMIN/
// parse
func (s *UserStore) ParseUpdateRole(r *http.Request) (*models.UserRolePayload, error) {
	var payload models.UserRolePayload
	if err := myhttp.ParseJSON(r, &payload); err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid request")
	}

	if payload.Role == "" {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "missing credentials")
	}

	return &payload, nil
}

// read
func (s *UserStore) GetUsersByRolePage(role string, page int) ([]models.ShowProfile, error) {
	if page < 1 {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid query")
	}

	limit := 10
	offset := (page - 1) * limit

	valideRoles := map[string]bool{
		"admin":    true,
		"employee": true,
		"user":     true,
		"all":      true,
	}

	if exists := valideRoles[role]; !exists {
		return nil, myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid role")
	}

	var query *gorm.DB
	if role == "all" {
		query = s.db.Limit(limit).Offset(offset)
	} else {
		query = s.db.Where("role = ?", role).Limit(limit).Offset(offset)
	}

	var users []models.User
	if err := query.Find(&users).Error; err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	var response []models.ShowProfile
	for _, user := range users {
		response = append(response, *user.ToShowProfile())
	}

	if page > 1 && len(response) == 0 {
		return nil, myhttp.NewErrorHTTP(http.StatusNotFound, "no users found")
	}

	return response, nil
}

func (s *UserStore) PatchRoleUser(id int, role string) error {
	if id < 1 {
		return myhttp.NewErrorHTTP(http.StatusBadRequest, "invalid id")
	}

	if err := s.CheckRole(role); err != nil {
		return err
	}

	if err := s.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]any{"role": role}).Error; err != nil {
		return myhttp.NewErrorHTTP(http.StatusInternalServerError, err.Error())
	}

	return nil
}
