package postdto

import (
	"social/internal/domain/post"
	"time"
)

type PostResponse struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}

func ToPostResponse(p *post.Post) *PostResponse {
	return &PostResponse{
		Id:          p.Id,
		UserId:      p.UserId,
		Title:       p.Title,
		Description: p.Description,
		Tags:        p.Tags,
		CreatedAt:   p.CreatedAt,
	}
}
