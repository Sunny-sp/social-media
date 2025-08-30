package user

import (
	"context"
)

type UserRepository interface {
	GetByUserId(ctx context.Context, userID int64) (*User, error)
	GetUserPassByUserId(ctx context.Context, userID int64) (*DBUser, error)
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *DBUser) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
