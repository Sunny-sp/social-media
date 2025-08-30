package repository

import (
	"context"
	"errors"
	"social/internal/domain/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) user.UserRepository {
	return &UserRepo{pool: pool}
}

func (r *UserRepo) GetByUserId(ctx context.Context, userID int64) (*user.User, error) {
	u := &user.User{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile FROM "user" WHERE user_id=$1`, userID).
		Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetUserPassByUserId(ctx context.Context, userID int64) (*user.DBUser, error) {
	u := &user.DBUser{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile, password FROM "user" WHERE user_id=$1`, userID).
		Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile, &u.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]*user.User, error) {
	rows, err := r.pool.Query(ctx, `SELECT user_id, name, email, mobile FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		if err := rows.Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Create(ctx context.Context, u *user.DBUser) (*user.User, error) {
	created := &user.User{}
	err := r.pool.QueryRow(ctx,
		`INSERT INTO "user" (name, email, mobile, password) VALUES ($1, $2, $3, $4) RETURNING user_id, name, email, mobile`,
		u.Name, u.Email, u.Mobile, u.Password).
		Scan(&created.UserId, &created.Name, &created.Email, &created.Mobile)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	u := &user.User{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile FROM "user" WHERE email=$1`, email).
		Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
