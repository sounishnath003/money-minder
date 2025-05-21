package server

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

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

	// Serve static files from the dist directory
	fs := http.FileServer(http.Dir("./web/dist/"))

	// Handle all routes by serving index.html for client-side routing
	mux.Handle("/", fs)

	// Enforce middleware chains on the mux with CORS middleware
	http.ListenAndServe(fmt.Sprintf(":%d", s.Co.Port),
		MiddlewareChain(
			cors.New(cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"GET", "POST", "HEAD"},
				ExposedHeaders: []string{"X-Request-ID", "Authorization", "Content-Type", "Content-Length"},
				MaxAge:         3600,
			}).Handler(mux),
			LoggerMiddleware(s.Co),
			TimeoutMiddleware(s.Co),
			AddContextMiddleware("co", s.Co)))
}

// defineApiRouterV1Endpoints - define a subrouter for the api groupping.
// Group of /api/v1 routes
func defineApiRouterV1Endpoints() http.Handler {
	router := http.NewServeMux()
	// API /api/v1 route endpoints
	router.HandleFunc("GET /healthz", v1.HealthHandler)
	router.HandleFunc("GET /categories", v1.GetAllCategoriesHandler)
	router.HandleFunc("GET /transactions", v1.GetTransactionsHandler)
	router.HandleFunc("POST /transactions", v1.AddTransactionHandler)
	router.HandleFunc("GET /transactions/spendByCategory", v1.GetTotalSpendByCategoryHandler)

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
