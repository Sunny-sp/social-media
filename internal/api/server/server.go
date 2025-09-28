package server

import (
	"fmt"
	"net/http"
	"social/internal/api/auth_api"
	mw "social/internal/api/middleware"
	"social/internal/api/post_api"
	"social/internal/api/user_api"
	"social/internal/config"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config         config.ServerConfig
	authMiddleware *mw.AuthMiddleware
	userHandler    *user_api.UserHandler
	authHandler    *auth_api.AuthHandler
	postHandler    *post_api.PostHandler
}

func NewServer(
	cfg config.ServerConfig,
	authMiddleware *mw.AuthMiddleware,
	userHandler *user_api.UserHandler,
	authHandler *auth_api.AuthHandler,
	postHandler *post_api.PostHandler,
) *Server {
	return &Server{
		config:         cfg,
		authMiddleware: authMiddleware,
		userHandler:    userHandler,
		authHandler:    authHandler,
		postHandler:    postHandler,
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
		user_api.UserRoutes(r, s.userHandler, s.authMiddleware)
		auth_api.AuthRoutes(r, s.authHandler)
		post_api.PostRoutes(r, s.postHandler, s.authMiddleware)
	})

	return r
}

func PrintRoutes(r chi.Router) {
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("📌 %s %s\n", method, route)
		return nil
	})
}

func (s *Server) Run(mux http.Handler) error {
	addr := s.config.Addr()

	fmt.Printf("🚀 Server is starting at port%s\n", addr)
	// PrintRoutes(mux.(chi.Router))

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return srv.ListenAndServe()
}
