package post_api

import (
	"social/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func PostRoutes(r chi.Router, postHandler *PostHandler, authMiddleware *middleware.AuthMiddleware) {
	r.Route("/posts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.RequireAuth)
			r.Post("/", postHandler.AddPost)
			r.Get("/{id}", postHandler.getPostById)
		})
	})
}
