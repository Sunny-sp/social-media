package views

import "time"

type PostView struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"Description"`
	UserId      int64     `json:"user_id"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}
