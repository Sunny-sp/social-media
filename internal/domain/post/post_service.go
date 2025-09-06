package post

import (
	"context"
	"fmt"
)

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (p *PostService) AddPost(ctx context.Context, post *Post) (*Post, error) {
	createdPost, err := p.postRepo.AddNewPost(ctx, post)

	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (p *PostService) GetPostById(ctx context.Context, id int64) (*Post, error) {
	post, err := p.postRepo.GetPostById(ctx, id)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

