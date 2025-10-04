package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func (r *UserRepo) GetProfileByUserId(ctx context.Context, userId int64) (*user.Profile, error) {
	u := &user.Profile{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile, profile_pic, bio FROM "user" WHERE user_id=$1`, userId).
		Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile, &u.ProfilePic, &u.Bio)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) UpdateProfileByUserId(ctx context.Context, userId int64, p *user.Profile) error {
	query := `update "user" set `
	args := []any{}
	i := 1

	if p.Name != "" {
		query += fmt.Sprintf("name = $%d,", i)
		args = append(args, p.Name)
		i++
	}

	if p.Email != "" {
		query += fmt.Sprintf("email = $%d,", i)
		args = append(args, p.Email)
		i++
	}
	if p.Mobile != "" {
		query += fmt.Sprintf("mobile = $%d,", i)
		args = append(args, p.Mobile)
		i++
	}
	if p.ProfilePic != nil {
		query += fmt.Sprintf("profile_pic = $%d,", i)
		args = append(args, *p.ProfilePic)
		i++
	}
	if p.Bio != nil {
		query += fmt.Sprintf("bio = $%d,", i)
		args = append(args, *p.Bio)
		i++
	}
	query += fmt.Sprintf("updated_at = now() where user_id = $%d", i)

	args = append(args, userId)
	log.Println("::::query::", query)
	log.Println("::::args::", args)
	_, err := r.pool.Exec(ctx, query, args...)
	return err
}

func (r *UserRepo) GetUserByUserId(ctx context.Context, userId int64) (*user.User, error) {
	u := &user.User{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile, profile_pic, bio FROM "user" WHERE user_id=$1`, userId).
		Scan(&u.UserId, &u.Name, &u.Email, &u.Mobile, &u.ProfilePic, &u.Bio)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetUserPassByUserId(ctx context.Context, userId int64) (*user.User, error) {
	u := &user.User{}
	err := r.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile, password FROM "user" WHERE user_id=$1`, userId).
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

func (r *UserRepo) Create(ctx context.Context, u *user.User) (*user.User, error) {
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
