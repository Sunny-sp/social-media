package userdto

import (
	"social/internal/domain/user"
	"time"
)

type UserResponse struct {
	UserId    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(u *user.User) *UserResponse {
	return &UserResponse{
		UserId:    u.UserId,
		Name:      u.Name,
		Email:     u.Email,
		Mobile:    u.Mobile,
		CreatedAt: u.CreatedAt,
	}
}
