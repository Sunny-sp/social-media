package service

import (
	"context"
	"fmt"
	"log"
	"social/internal/api/model"
	"social/internal/api/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserByUserId(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.repo.GetByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetAllUser(ctx context.Context) ([]*model.User, error) {
	users, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserService) CreateNewUser(ctx context.Context, user *model.DBUser) (*model.User, error) {
	existing, err := u.repo.GetByEmail(ctx, user.Email)

	log.Println("existing", existing)

	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}

	createdUser, err := u.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
