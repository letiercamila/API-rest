package middleware

import "net/http"

func SetContentType(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("Content-Type", "application/json")
		nxt.ServeHTTP(wr, r)
	})
}
