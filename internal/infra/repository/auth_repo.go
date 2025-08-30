package repository

import (
	"context"
	"errors"
	"social/internal/domain/auth"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	pool *pgxpool.Pool
}

func NewAuthRepo(pool *pgxpool.Pool) auth.AuthRepository {
	return &AuthRepo{pool: pool}
}

func (r *AuthRepo) GetTokenByUserID(ctx context.Context, userID int64) (*auth.AuthToken, error) {
	token := &auth.AuthToken{}
	err := r.pool.QueryRow(ctx,
		`SELECT token_id, user_id, token, expires_at FROM auth_tokens WHERE user_id=$1`,
		userID).Scan(&token.TokenID, &token.UserID, &token.Token, &token.ExpiresAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Not found
		}
		return nil, err
	}
	return token, nil
}
