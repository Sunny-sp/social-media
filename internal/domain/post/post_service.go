package post

import (
	"context"
	"fmt"
	"path/filepath"
	"social/internal/storage"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo   PostRepository
	storageSvc storage.StorageService
}

func NewPostService(postRepo PostRepository, storageSrv storage.StorageService) *PostService {
	return &PostService{
		postRepo:   postRepo,
		storageSvc: storageSrv,
	}
}

func (p *PostService) AddPost(ctx context.Context, post *Post) error {
	err := p.postRepo.AddNewPost(ctx, post)

	if err != nil {
		return err
	}

	return nil
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

func (p *PostService) GeneratePostSignedURL(ctx context.Context, originalFilename string, userId int64, expires time.Duration) (*PostPresignedResult, error) {
	var postDir = "social-posts"
	ext := filepath.Ext(originalFilename)
	key := fmt.Sprintf("%s/%d/%s%s", postDir, userId, uuid.New().String(), ext)
	signedURL, err := p.storageSvc.GeneratePresignedPutURL(ctx, key, expires)
	if err != nil {
		return nil, fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return &PostPresignedResult{
		Key:              key,
		OriginalFilename: originalFilename,
		SignedURL:        signedURL,
	}, nil
}
