package store

import (
	"context"
	"net/http"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/internal/utils/auth"
	"github.com/Ayeye11/inv/pkg/myhttp"
	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareStore struct {
	jwtKey string
}

var valideRoles = map[string]int{
	"user":     1,
	"employee": 2,
	"admin":    3,
}

// auth
func (s *MiddlewareStore) GetClaimsFromCookie(r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusUnauthorized, "missing cookie")
	}

	if cookie.Value == "" {
		return nil, myhttp.NewErrorHTTP(http.StatusUnauthorized, "missing token")
	}

	claims, err := auth.CheckToken(cookie.Value, s.jwtKey)
	if err != nil {
		return nil, myhttp.NewErrorHTTP(http.StatusUnauthorized, err.Error())
	}

	return claims, nil
}

func (s *MiddlewareStore) GetSingleClaim(claims jwt.MapClaims, key string) (string, error) {
	if claim, ok := claims[key].(string); ok {
		return claim, nil
	}

	return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "failed to get claim")
}

func (s *MiddlewareStore) CheckRole(role, minRole string) error {
	val, exists := valideRoles[minRole]
	if !exists {
		return myhttp.NewErrorHTTP(http.StatusInternalServerError, "invalid user role")
	}

	userRole, exists := valideRoles[role]
	if !exists {
		return myhttp.NewErrorHTTP(http.StatusInternalServerError, "invalid user role")
	}

	if val >= userRole {
		return nil
	}

	return myhttp.NewErrorHTTP(http.StatusForbidden, "forbidden")
}

// context
func (s *MiddlewareStore) SetClaimsToContext(r *http.Request, claims jwt.MapClaims) context.Context {
	var ctxKey models.ContextKey = "claims"
	return context.WithValue(r.Context(), ctxKey, claims)
}
