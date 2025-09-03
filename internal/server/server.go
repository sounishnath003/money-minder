package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"

	"github.com/sounishnath003/money-minder/internal/core"
	v1 "github.com/sounishnath003/money-minder/internal/server/handlers/v1"
)

type Server struct {
	Co *core.Core
}

// getAllowedOrigins returns a list of allowed origins for CORS.
// In production, you should replace this with a config/env variable or a more secure mechanism.
func getAllowedOrigins() []string {
	// Example: allow localhost and your production domain
	return []string{
		"http://localhost:5173",
		"http://localhost:3000",
	}
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

	// --- CORS FIX START ---
	// Use github.com/rs/cors as a top-level handler, not inside MiddlewareChain.
	// This ensures CORS headers are set for all requests, including OPTIONS preflight.
	// AllowCredentials cannot be used with AllowedOrigins: ["*"], so we must specify explicit origins.
	allowedOrigins := getAllowedOrigins()
	c := cors.New(cors.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedMethods:     []string{"GET", "POST", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"X-Request-ID", "Authorization", "Content-Type", "Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials:   true,
		MaxAge:             3600,
		OptionsPassthrough: false,
		// Enable origin function for more dynamic control if needed
		AllowOriginFunc: func(origin string) bool {
			for _, o := range allowedOrigins {
				if o == origin {
					return true
				}
				// Optionally allow subdomains
				if strings.HasPrefix(o, "https://") && strings.HasPrefix(origin, "https://") {
					if strings.HasSuffix(origin, strings.TrimPrefix(o, "https://")) {
						return true
					}
				}
			}
			return false
		},
	})
	handlerWithMiddleware := MiddlewareChain(
		mux,
		LoggerMiddleware(s.Co),
		TimeoutMiddleware(s.Co),
		AddContextMiddleware("co", s.Co),
	)
	finalHandler := c.Handler(handlerWithMiddleware)

	// Listen and serve with the CORS-wrapped handler
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Co.Port), finalHandler); err != nil {
		fmt.Fprintf(os.Stderr, "server failed: %v\n", err)
	}
}

// defineApiRouterV1Endpoints - define a subrouter for the api groupping.
// Group of /api/v1 routes
func defineApiRouterV1Endpoints() http.Handler {
	router := http.NewServeMux()
	// Auth endpoints (public)
	router.HandleFunc("POST /auth/register", v1.RegisterHandler)
	router.HandleFunc("POST /auth/login", v1.LoginHandler)
	// Public endpoints
	router.HandleFunc("GET /healthz", v1.HealthHandler)
	router.HandleFunc("GET /categories", v1.GetAllCategoriesHandler)
	router.HandleFunc("GET /paymentMethods", v1.GetAllPaymentMethodOptionsHandler)

	// Protected endpoints (require JWT)
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("GET /transactions", v1.GetTransactionsHandler)
	protectedMux.HandleFunc("POST /transactions", v1.AddTransactionHandler)
	protectedMux.HandleFunc("GET /transactions/spendByCategory", v1.GetTotalSpendByCategoryHandler)
	protectedMux.HandleFunc("GET /analytics/getDailySpendByTimeframe", v1.GetDailyTotalSpendByTimeframeHandler)
	protectedMux.HandleFunc("GET /analytics/getSpendOnCategoriesMonthOnMonth", v1.GetSpendOnCategoriesMonthOnMonthHandler)
	protectedMux.HandleFunc("GET /analytics/getSpendOnCategoriesByAggregatedByYearMonth", v1.GetSpendOnCategoriesByAllYearMonthAggregatedHandler)
	protectedMux.HandleFunc("GET /analytics/getMonthlySavingsBreakdown", v1.GetMonthlySavingsBreakdownHandler)

	router.Handle("/transactions", MiddlewareChain(protectedMux, AuthJWTMiddleware()))
	router.Handle("/transactions/", MiddlewareChain(protectedMux, AuthJWTMiddleware()))
	router.Handle("/analytics/", MiddlewareChain(protectedMux, AuthJWTMiddleware()))

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
