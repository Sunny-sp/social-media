package userdto

import (
	"social/internal/domain/user"
	"time"
)

type UserResponse struct {
	UserId     int64     `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	ProfilePic string    `json:"profile_pic"`
	Bio        string    `json:"bio"`
	CreatedAt  time.Time `json:"created_at"`
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

func ToProfileResponse(u *user.Profile) *UserResponse {
	var bio string
	if u.Bio != nil {
		bio = *u.Bio
	}

	return &UserResponse{
		UserId:    u.UserId,
		Name:      u.Name,
		Email:     u.Email,
		Mobile:    u.Mobile,
		Bio:       bio,
		CreatedAt: u.CreatedAt,
	}
}
