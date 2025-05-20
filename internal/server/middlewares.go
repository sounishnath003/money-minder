package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sounishnath003/money-minder/internal/core"
)

type Middleware func(http.Handler) http.Handler

// AddContextMiddleware helps to add context into the requests based on the key provided
func AddContextMiddleware(key string, co *core.Core) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a new context with the core instance
			ctx := context.WithValue(r.Context(), key, co)

			// Create a new request with the updated context
			r = r.WithContext(ctx)
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// AddCorsMiddleware helps to add domain as middleware CORS allowed domain
func AddCorsMiddleware(domain string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", domain)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// LoggerMiddleware helps to check the logger middleware, which logs
// all the requests which ever comes to the server
func LoggerMiddleware(co *core.Core) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log request details
			co.Logger.Printf("RemoteIp: %s UserAgent: %s Request: %s %s", r.RemoteAddr, r.UserAgent(), r.Method, r.URL.Path)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// TimeoutMiddleware - helps to track the request ttl.
// The request will get auto timeout if it takes for than 5 seconds
func TimeoutMiddleware(co *core.Core) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a context with timeout
			ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			// Create a channel to signal completion
			done := make(chan struct{})

			// Create a new request with the timeout context
			r = r.WithContext(ctx)

			// Handle the request in a goroutine
			go func() {
				next.ServeHTTP(w, r)
				close(done)
			}()

			// Wait for either completion or timeout
			select {
			case <-done:
				// Request completed successfully
				return
			case <-ctx.Done():
				// Request timed out
				http.Error(w, "Request timeout", http.StatusGatewayTimeout)
				return
			}
		})
	}
}

func MiddlewareChain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
