package routes

import (
	"social/internal/api/controller"

	"github.com/go-chi/chi/v5"
)

func RegistorRoutes(r *chi.Mux, ctrl *controller.Controller) {
	r.Route("/v1", func(r chi.Router) {
		UserRouter(r, ctrl.UserControlller)
		AuthRouter(r, ctrl.AuthController)
	})
}
