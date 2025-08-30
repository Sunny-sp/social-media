package auth

import (
	"context"
)

type AuthRepository interface {
	GetTokenByUserID(ctx context.Context, userID int64) (*AuthToken, error)
}
