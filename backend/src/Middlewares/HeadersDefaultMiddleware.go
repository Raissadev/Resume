package middlewares

import (
	"net/http"
)

type HeadersDefaultMiddleware struct {
	//
}

func (pmw *HeadersDefaultMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
		return
	})
}
