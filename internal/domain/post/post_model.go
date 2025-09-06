package post

import "time"

type Post struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
