package postdto

import "social/internal/domain/post"

type CreatePostDto struct {
	UserId      int64
	Title       string
	Description string
	Tags        []string
}

func (p *CreatePostDto) ToDomain() *post.Post {
	return &post.Post{
		UserId:      p.UserId,
		Title:       p.Title,
		Description: p.Description,
		Tags:        p.Tags,
	}
}
