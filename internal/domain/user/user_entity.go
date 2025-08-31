package user

import "time"

type User struct {
	UserId    int64
	Name      string
	Email     string
	Mobile    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
