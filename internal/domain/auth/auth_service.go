package auth

import (
	"context"
	"fmt"
	authdto "social/internal/domain/auth/dtos"
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

func (s *AuthService) ValidateUser(ctx context.Context, loginDto *authdto.LoginDto) (*user.User, string, error) {
	dbUser, err := s.userRepo.GetUserPassByUserId(ctx, loginDto.UserId)
	if err != nil {
		return nil, "", fmt.Errorf("invalid user id or password: %w", err)
	}
	if dbUser == nil {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	// TODO: Implement password decryption/hashing comparison
	if dbUser.Password != loginDto.Password {
		return nil, "", fmt.Errorf("invalid user id or password")
	}

	token, err := utils.GenrateJWT(dbUser.UserId, dbUser.Email, s.jwtSecret, s.jwtExpiry)

	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return &user.User{
		UserId: dbUser.UserId,
		Name:   dbUser.Name,
		Email:  dbUser.Email,
		Mobile: dbUser.Mobile,
	}, token, nil
}
