package service

import (
	"context"
	"fmt"
	"log"
	dto "social/internal/api/dto/auth"
	"social/internal/api/model"
	"social/internal/api/repository"
	"social/internal/api/utils"
)

type AuthService struct {
	// authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

func NewAuthService(
	// authRepo repository.AuthRepository,
	userRepo repository.UserRepository,
) *AuthService {
	return &AuthService{
		// authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *AuthService) ValidateUser(ctx context.Context, loginDto *dto.LoginDto) (*model.User, string, error) {

	user, err := s.userRepo.GetUserPassByUserId(ctx, loginDto.UserId)
	log.Println("User:", user)
	log.Println("err:", err)

	if err != nil {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	// TODO Later: decrypt user pass and match with userPass

	if user.Password != loginDto.Password {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	token, err := utils.GenrateJWT(user.UserId, user.Email, []byte("secret"))

	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return &model.User{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
	}, token, nil
}
