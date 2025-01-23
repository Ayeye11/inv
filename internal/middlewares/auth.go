package middlewares

import (
	"net/http"
)

func (m *Middleware) AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
