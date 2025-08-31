package auth_api

import (
	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, authHandler *AuthHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})
}
