package adapters

import (
	"context"
	"social/internal/domain/post"
	"social/internal/domain/user"
	"social/internal/domain/user/views"
	"strings"
)

type PostProviderAdapter struct {
	postRepo post.PostRepository
}

func NewPostProviderAdapter(postRepo post.PostRepository) *PostProviderAdapter {
	return &PostProviderAdapter{
		postRepo: postRepo,
	}
}

func (p *PostProviderAdapter) GetPostsByUserId(ctx context.Context, UserId int64, f user.PostFilter) ([]*views.PostView, error) {
	// convert string tags into []string
	var tags []string
	if f.Filters.Tags != "" {
		tags = strings.Split(strings.TrimSpace(f.Filters.Tags), ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
	}
	query := post.PostByUserIdFilter{
		Page:   f.Page,
		Limit:  f.Limit,
		Search: f.Search,
		Sort:   f.Sort,
		Filters: post.PostFilters{
			Title: f.Filters.Title,
			Tags:  tags,
			Desc:  f.Filters.Desc,
		},
	}

	posts, err := p.postRepo.GetPostsByUserId(ctx, UserId, query)

	if err != nil {
		return nil, err
	}

	postViews := make([]*views.PostView, len(posts))

	for i, p := range posts {

		mediaViews := make([]views.Media, len(p.MediaURLs))

		for i, m := range p.MediaURLs {
			mediaViews[i] = views.Media{
				Type: m.Type,
				Path: m.Path,
			}
		}

		postViews[i] = &views.PostView{
			Id:          p.Id,
			UserId:      p.UserId,
			Title:       p.Title,
			Description: p.Description,
			Tags:        p.Tags,
			MediaURLs:   mediaViews,
			CreatedAt:   p.CreatedAt,
		}
	}

	return postViews, nil
}
