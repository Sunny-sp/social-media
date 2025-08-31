package post

import "context"

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (p *PostService) AddPost(ctx context.Context, post *Post) {
}
