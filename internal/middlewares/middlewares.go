package middlewares

import "github.com/Ayeye11/inv/internal/store"

type Middleware struct {
	store store.MiddlewareRepository
}

func SetMiddlewares(store store.MiddlewareRepository) *Middleware {
	return &Middleware{store}
}
