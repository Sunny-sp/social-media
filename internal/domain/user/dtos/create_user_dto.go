package userdto

import "social/internal/domain/user"

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (dto *CreateUserDTO) ToModel() *user.DBUser {
	return &user.DBUser{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
