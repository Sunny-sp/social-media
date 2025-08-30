package routes

import (
	"social/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, userHandler *handlers.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAllUser)
		r.Post("/", userHandler.Create)
		r.Get("/{id}", userHandler.GetByUserId)
	})
}
