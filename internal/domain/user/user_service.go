package user

import (
	"context"
	"fmt"
	"log"
	"social/internal/domain/user/views"
)

type UserService struct {
	userRepo     UserRepository
	postProvider PostProvider
}

func NewUserService(userRepo UserRepository, postProvider PostProvider) *UserService {
	return &UserService{
		userRepo:     userRepo,
		postProvider: postProvider,
	}
}

func (u *UserService) GetUserByUserId(ctx context.Context, id int64) (*User, error) {
	user, err := u.userRepo.GetByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
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

func (u *UserService) GetPostsByUserId(ctx context.Context, UserId int64) ([]*views.PostView, error) {
	//check if user exist otherwise error
	user, err := u.GetUserByUserId(ctx, UserId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}


	return u.postProvider.GetPostsByUserId(ctx, UserId)
}
