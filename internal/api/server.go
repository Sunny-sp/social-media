package api

import (
	"net/http"
	"social/internal/api/handlers"
	"social/internal/api/routes"
	"social/internal/infra/config"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config      config.ServerConfig
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
}

func NewServer(cfg config.ServerConfig, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) *Server {
	return &Server{
		config:      cfg,
		userHandler: userHandler,
		authHandler: authHandler,
	}
}

func (s *Server) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		routes.UserRoutes(r, s.userHandler)
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
