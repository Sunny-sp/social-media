package user

import (
	"context"
	"fmt"
	"log"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetUserByUserId(ctx context.Context, id int64) (*User, error) {
	user, err := u.userRepo.GetByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetAllUser(ctx context.Context) ([]*User, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) CreateNewUser(ctx context.Context, dto *User) (*User, error) {
	existing, err := u.userRepo.GetByEmail(ctx, dto.Email)

	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}
	user := &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Mobile:   dto.Mobile,
	}

	createdUser, err := u.userRepo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
