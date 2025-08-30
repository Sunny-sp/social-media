package controller

import "social/internal/api/service"

type Controller struct {
	UserControlller *UserController
	AuthController  *AuthController
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		UserControlller: NewUserController(s.UserService),
		AuthController:  NewAuthController(s.AuthService),
	}
}
