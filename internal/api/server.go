package api

import (
	"net/http"
	"social/internal/api/handlers"
	mw "social/internal/api/middleware"
	"social/internal/api/routes"
	"social/internal/infra/config"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config         config.ServerConfig
	authMiddleware *mw.AuthMiddleware
	userHandler    *handlers.UserHandler
	authHandler    *handlers.AuthHandler
}

func NewServer(
	cfg config.ServerConfig,
	authMiddleware *mw.AuthMiddleware,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
) *Server {
	return &Server{
		config:         cfg,
		authMiddleware: authMiddleware,
		userHandler:    userHandler,
		authHandler:    authHandler,
	}
}

func (s *Server) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		routes.UserRoutes(r, s.userHandler, s.authMiddleware)
		routes.AuthRoutes(r, s.authHandler)
	})

	return r
}

func (s *Server) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         s.config.Addr(),
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return srv.ListenAndServe()
}
