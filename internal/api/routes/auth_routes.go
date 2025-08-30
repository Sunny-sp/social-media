package routes

import (
	"social/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, authHandler *handlers.AuthHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})
}
