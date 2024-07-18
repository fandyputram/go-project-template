package middleware

import "net/http"

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement logging
		next.ServeHTTP(w, r)
	})
}
