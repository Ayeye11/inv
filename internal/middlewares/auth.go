package middlewares

import (
	"net/http"

	"github.com/Ayeye11/inv/pkg/myhttp"
)

func (m *Middleware) AuthWithClaims(next http.Handler) http.Handler {
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

func (m *Middleware) AuthEmployeeWithClaims(next http.Handler) http.Handler {
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

func (m *Middleware) AuthAdminWithClaims(next http.Handler) http.Handler {
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

// without claims
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := m.store.GetClaimsFromCookie(r)
		if err != nil {
			myhttp.SendError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) AuthEmployee(next http.Handler) http.Handler {
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

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) AuthAdmin(next http.Handler) http.Handler {
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

		next.ServeHTTP(w, r)
	})
}
