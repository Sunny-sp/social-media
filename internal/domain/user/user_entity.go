package user

import "time"

type User struct {
	UserId     int64
	Name       string
	Email      string
	Mobile     string
	Password   string
	ProfilePic *string
	Bio        *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Profile struct {
	UserId     int64
	Name       string
	Email      string
	Mobile     string
	ProfilePic *string
	Bio        *string
	CreatedAt  time.Time
}
