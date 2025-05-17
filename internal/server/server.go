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

	go func() {
		s.Co.Logger.Printf("server has been started on http://%s:%d", s.Co.Hostname, s.Co.Port)
	}()
	// Enforce middleware chains on the mux
	http.ListenAndServe(fmt.Sprintf(":%d", s.Co.Port), MiddlewareChain(mux, LoggerMiddleware(s.Co), TimeoutMiddleware(s.Co)))
}

// defineApiRouterV1Endpoints - define a subrouter for the api groupping.
// Group of /api/v1 routes
func defineApiRouterV1Endpoints() http.Handler {
	routeV1 := http.NewServeMux()
	// API /api/v1 route endpoints
	routeV1.HandleFunc("GET /healthz", v1.HealthHandler)

	return routeV1
}
