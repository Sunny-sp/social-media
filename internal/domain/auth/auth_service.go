package auth

import (
	"context"
	"fmt"
	"social/internal/domain/user"
	"social/internal/pkg/utils"
)

type AuthService struct {
	authRepo  AuthRepository
	userRepo  user.UserRepository
	jwtSecret []byte
	jwtExpiry int
}

func NewAuthService(authRepo AuthRepository, userRepo user.UserRepository, jwtSecret []byte, jwtExpiry int) *AuthService {

	return &AuthService{
		authRepo:  authRepo,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
	}
}

func (s *AuthService) ValidateUser(ctx context.Context, loginDto *LoginCredentials) (*user.User, string, error) {
	user, err := s.userRepo.GetUserPassByUserId(ctx, loginDto.UserId)

	if err != nil || user == nil {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	// TODO: Implement password decryption/hashing comparison
	if user.Password != loginDto.Password {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	token, err := utils.GenrateJWT(user.UserId, user.Email, s.jwtSecret, s.jwtExpiry)

	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	user.Password = ""
	return user, token, nil
}
