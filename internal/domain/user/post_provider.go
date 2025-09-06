// internal/domain/user/post_provider.go
package user

import (
    "context"
    "social/internal/domain/user/views"
)

type PostProvider interface {
    GetPostsByUserId(ctx context.Context, userId int64) ([]*views.PostView, error)
}
