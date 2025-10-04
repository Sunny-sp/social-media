package user

import (
	"context"
)

type UserRepository interface {
	GetProfileByUserId(ctx context.Context, userId int64) (*Profile, error)
	UpdateProfileByUserId(ctx context.Context, userId int64, p *Profile) error
	GetUserByUserId(ctx context.Context, userId int64) (*User, error)
	GetUserPassByUserId(ctx context.Context, userId int64) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
