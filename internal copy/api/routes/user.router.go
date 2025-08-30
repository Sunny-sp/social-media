package routes

import (
	"social/internal/api/controller"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router, userCtrl *controller.UserController) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userCtrl.GetAllUser)
		r.Post("/", userCtrl.Create)
		r.Get("/{id}", userCtrl.GetByUserId)
	})
}
