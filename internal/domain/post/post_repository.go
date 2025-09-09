package post

import "context"

type PostRepository interface {
	AddNewPost(ctx context.Context, postData *Post) (error)
	GetPostById(ctx context.Context, id int64) (*Post, error)
	GetPostsByUserId(ctx context.Context, userId int64) ([]*Post, error)
}
