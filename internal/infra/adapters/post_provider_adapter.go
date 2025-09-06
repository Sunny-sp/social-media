package adapters

import (
	"context"
	"social/internal/domain/post"
	"social/internal/domain/user/views"
)

type PostProviderAdapter struct {
	postRepo post.PostRepository
}

func NewPostProviderAdapter(postRepo post.PostRepository) *PostProviderAdapter {
	return &PostProviderAdapter{
		postRepo: postRepo,
	}
}

func (p *PostProviderAdapter) GetPostsByUserId(ctx context.Context, UserId int64) ([]*views.PostView, error) {
	posts, err := p.postRepo.GetPostsByUserId(ctx, UserId)

	if err != nil {
		return nil, err
	}

	postViews := make([]*views.PostView, len(posts))

	for i, p := range posts {
		postViews[i] = &views.PostView{
			Id:          p.Id,
			UserId:      p.UserId,
			Title:       p.Title,
			Description: p.Description,
			Tags:        p.Tags,
			CreatedAt:   p.CreatedAt,
		}
	}

	return postViews, nil
}
