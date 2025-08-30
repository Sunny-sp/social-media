package user

import "time"

type User struct {
	UserId    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DBUser struct {
	UserId    int64     `db:"user_id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Mobile    string    `db:"mobile"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
