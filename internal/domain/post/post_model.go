package post

import "time"

type Media struct {
	Type string
	Path string
}

type Post struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	Tags        []string
	MediaURLs   []Media
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
