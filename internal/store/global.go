package store

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ayeye11/inv/internal/db/models"
	"github.com/Ayeye11/inv/pkg/myhttp"
	"github.com/golang-jwt/jwt/v5"
)

type GlobalStore struct {
}

// parse
func (s *GlobalStore) Atoi(x string) (int, error) {
	num, err := strconv.Atoi(x)
	if err != nil {
		return 0, myhttp.NewErrorHTTP(http.StatusBadRequest, fmt.Sprintf("invalid input: '%s'", x))
	}

	return num, nil
}

// context
func (s *GlobalStore) GetClaimsFromContext(r *http.Request) (jwt.MapClaims, error) {
	var ctx models.ContextKey = "claims"

	if claims, ok := r.Context().Value(ctx).(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail to get claims from context")
}

func (s *GlobalStore) GetSingleClaimFromContext(r *http.Request, key string) (string, error) {
	var ctx models.ContextKey = "claims"

	claims, ok := r.Context().Value(ctx).(jwt.MapClaims)
	if !ok {
		return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail to get claims from context")
	}

	if claim, ok := claims[key].(string); ok {
		return claim, nil
	}

	return "", myhttp.NewErrorHTTP(http.StatusInternalServerError, "fail to parse claims")
}
