package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sounishnath003/money-minder/internal/core"
)

type Middleware func(http.Handler) http.Handler

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
