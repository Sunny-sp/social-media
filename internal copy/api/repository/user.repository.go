package repository

import (
	"context"
	"errors"
	"log"
	"social/internal/api/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetByUserId(ctx context.Context, user_id int64) (*model.User, error)
	GetUserPassByUserId(ctx context.Context, user_id int64) (*model.DBUser, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.DBUser) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{
		pool: pool,
	}
}

func (u *userRepository) GetByUserId(ctx context.Context, user_id int64) (*model.User, error) {
	row := u.pool.QueryRow(ctx, `SELECT "user_id", name, email, mobile FROM "user" WHERE user_id=$1`, user_id)

	user := &model.User{}

	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Mobile)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetUserPassByUserId(ctx context.Context, user_id int64) (*model.DBUser, error) {
	row := u.pool.QueryRow(ctx, `SELECT "user_id", name, email, mobile, password FROM "user" WHERE user_id=$1`, user_id)

	user := &model.DBUser{}

	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Mobile, &user.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	rows, err := u.pool.Query(ctx, `SELECT user_id, name, email, mobile FROM "user"`)
	if err != nil {
		log.Println("GetAll error:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Mobile)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// check for errors from iteration
	if err = rows.Err(); err != nil {
		log.Println("GetAll error:", err)
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Create(ctx context.Context, user *model.DBUser) (*model.User, error) {
	err := u.pool.QueryRow(ctx,
		`INSERT INTO "user" (name, email, mobile, password)
		 VALUES ($1, $2, $3, $4)
		 RETURNING user_id`,

		user.Name, user.Email, user.Mobile, user.Password,
	).Scan(&user.UserId)

	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	return &model.User{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
	}, nil

}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	row := u.pool.QueryRow(ctx, `SELECT user_id, name, email, mobile FROM "user" WHERE email=$1`, email)

	user := &model.User{}
	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Mobile)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // no user found
		}
		return nil, err
	}
	return user, nil
}
