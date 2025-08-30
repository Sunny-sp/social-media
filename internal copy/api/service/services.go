package service

import "social/internal/api/repository"

type Service struct {
	UserService *UserService
	AuthService *AuthService
}

func NewService(r repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(r.UserRepo()),
		AuthService: NewAuthService(r.UserRepo()),
	}
}
