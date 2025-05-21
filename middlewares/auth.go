package middlewares

import "net/http"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("auth")
		if token != "secret" {
			http.Error(w, "Status Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
