package user_api

import (
	"social/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, userHandler *UserHandler, authMiddelware *middleware.AuthMiddleware) {
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authMiddelware.RequireAuth)
			r.Get("/", userHandler.GetAllUser)
			r.Get("/{id}", userHandler.GetByUserId)
		})
		r.Post("/", userHandler.Create)
	})
}
