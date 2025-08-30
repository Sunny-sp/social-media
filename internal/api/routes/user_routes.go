package routes

import (
	"social/internal/api/handlers"
	"social/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, userHandler *handlers.UserHandler, authMiddelware *middleware.AuthMiddleware) {
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authMiddelware.RequireAuth)
			r.Get("/", userHandler.GetAllUser)
			r.Get("/{id}", userHandler.GetByUserId)
		})
		r.Post("/", userHandler.Create)
	})
}
