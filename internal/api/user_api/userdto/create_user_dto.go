package userdto

import "social/internal/domain/user"

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Mobile   string `json:"mobile,omitempty"`
}

// its ok to import from domain into outside layers like haddler or infra/repository
// but ever import dtos into service this break seperations
func (dto *CreateUserDTO) ToDomain() *user.User {
	return &user.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Mobile:   dto.Mobile,
		Password: dto.Password,
	}
}
