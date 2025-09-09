package postdto

import (
	"social/internal/domain/post"
)

type MediaResponse struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type PostResponse struct {
	Id          int64           `json:"id"`
	UserId      int64           `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Tags        []string        `json:"tags"`
	MediaURLs   []MediaResponse `json:"media_urls"`
	CreatedAt   int64           `json:"created_at"`
}

func ToPostResponse(p *post.Post) *PostResponse {
	media := make([]MediaResponse, len(p.MediaURLs))
	for i, m := range p.MediaURLs {
		media[i] = MediaResponse{
			Type: m.Type,
			Path: m.Path,
		}
	}
	return &PostResponse{
		Id:          p.Id,
		UserId:      p.UserId,
		Title:       p.Title,
		Description: p.Description,
		Tags:        p.Tags,
		MediaURLs:   media,
		CreatedAt:   p.CreatedAt.Unix(),
	}
}
