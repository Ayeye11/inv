package middlewares

import (
	"net/http"

	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (m *Middleware) AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, err := m.store.GetClaimsFromCookie(r)
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		ctx := m.store.SetClaimsToContext(r, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) AuthorizeEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.store.GetClaimsFromCookie(r)
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		role, err := m.store.GetSingleClaim(claims, "role")
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		if err := m.store.CheckRole(role, "employee"); err != nil {
			myhttp.SendError(w, err)
			return
		}

		ctx := m.store.SetClaimsToContext(r, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, err := m.store.GetClaimsFromCookie(r)
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		role, err := m.store.GetSingleClaim(claims, "role")
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		if err := m.store.CheckRole(role, "admin"); err != nil {
			myhttp.SendError(w, err)
			return
		}

		ctx := m.store.SetClaimsToContext(r, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
