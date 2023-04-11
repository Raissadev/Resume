package middlewares

import (
	. "api/src/utils"
	"net/http"
	"os"
)

var env = Lenv

type PermissionMiddleware struct {
	Token map[string]string
}

func (pmw *PermissionMiddleware) Populate() {
	pmw.Token["key"] = os.Getenv("KEY_VALUE")
}

func (pmw *PermissionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if pmw.Token["key"] == token {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
