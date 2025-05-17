package server

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func LoggerMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log request details
			log.Printf("Request: %s %s", r.Method, r.URL.Path)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func MiddlewareChain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
