package server

import (
	"fmt"
	"net/http"

	"github.com/sounishnath003/money-minder/internal/core"
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

	// Add the api endpoints groups here
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", defineApiRouterV1Endpoints()))

	go func() {
		s.Co.Logger.Printf("server has been started on http://%s:%d", s.Co.Hostname, s.Co.Port)
	}()
	// Enforce middleware chains on the mux

	http.ListenAndServe(fmt.Sprintf(":%d", s.Co.Port), MiddlewareChain(mux, LoggerMiddleware()))
}

func defineApiRouterV1Endpoints() http.Handler {
	// Define a subrouter for the api groupping
	// Group of /api/v1 routes
	routeV1 := http.NewServeMux()
	// API /api/v1 route endpoints
	routeV1.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("server is up and running"))
	})

	return routeV1
}
