package routes

import (
	"social/internal/api/controller"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(r chi.Router, authCtr *controller.AuthController) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authCtr.Login)
	})
}
