package postdto

import "social/internal/domain/post"

type MediaDto struct {
	Type string `json:"type" validate:"required,oneof=image video"`
	Path string `json:"path" validate:"required"`
}

type CreatePostDto struct {
	UserId      int64      `json:"user_id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Tags        []string   `json:"tags,omitempty"`
	MediaURLs   []MediaDto `json:"media_urls" validate:"required,min=1,dive"` // must have at least 1
}

func (p *CreatePostDto) ToDomain() *post.Post {
	media := make([]post.Media, len(p.MediaURLs))
	for i, m := range p.MediaURLs {
		media[i] = post.Media{
			Type: m.Type,
			Path: m.Path,
		}
	}

	return &post.Post{
		UserId:      p.UserId,
		Title:       p.Title,
		Description: p.Description,
		Tags:        p.Tags,
		MediaURLs:   media,
	}
}
