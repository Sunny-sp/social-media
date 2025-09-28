// internal/domain/user/post_provider.go
package user

import (
	"context"
	"social/internal/domain/user/views"
)

type PostFilter struct {
	Page   int64
	Limit  int64
	Search string
	Sort   string

	Filters struct {
		Title string
		Desc  string
		Tags  string
	}
}

type PostProvider interface {
	GetPostsByUserId(ctx context.Context, userId int64, filter PostFilter) ([]*views.PostView, error)
}
