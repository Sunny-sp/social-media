package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	UserRepo() UserRepository
	// AuthRepo() AuthRepository
}

type repository struct {
	userRepo UserRepository
	// authRepo AuthRepository
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &repository{
		userRepo: NewUserRepository(pool),
		// authRepo: NewAuthRepository(pool),
	}
}

func (r *repository) UserRepo() UserRepository {
	return r.userRepo
}

// func (r *repository) AuthRepo() AuthRepository {
// 	return r.authRepo
// }
