package server

import (
	"fmt"
	"net/http"

	"github.com/sounishnath003/money-minder/internal/core"
	v1 "github.com/sounishnath003/money-minder/internal/server/handlers/v1"
)

type Server struct {
	Co *core.Core
}

func NewServer(co *core.Core) *Server {
	return &Server{
		Co: co,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	// Add the api/v1/* endpoints groups here
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", defineApiRouterV1Endpoints()))
	// Add the api/v2/* endpoints groups here
	mux.Handle("/api/v2/", http.StripPrefix("/api/v2", defineApiRouterV2Endpoints()))

	go func() {
		s.Co.Logger.Printf("server has been started on http://%s:%d", s.Co.Hostname, s.Co.Port)
	}()
	// Enforce middleware chains on the mux
	http.ListenAndServe(fmt.Sprintf(":%d", s.Co.Port), MiddlewareChain(mux, LoggerMiddleware(s.Co), TimeoutMiddleware(s.Co), AddContextMiddleware("co", s.Co)))
}

// defineApiRouterV1Endpoints - define a subrouter for the api groupping.
// Group of /api/v1 routes
func defineApiRouterV1Endpoints() http.Handler {
	router := http.NewServeMux()
	// API /api/v1 route endpoints
	router.HandleFunc("GET /healthz", v1.HealthHandler)
	router.HandleFunc("GET /transactions", v1.GetTransactionsHandler)

	return router
}

// defineApiRouterV2Endpoints - define a subrouter for the api groupping.
// Group of /api/v2 routes
func defineApiRouterV2Endpoints() http.Handler {
	router := http.NewServeMux()
	// API /api/v1 route endpoints
	router.HandleFunc("GET /healthz", v1.HealthHandler)

	return router
}
