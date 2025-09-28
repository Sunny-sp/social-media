package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Storage struct {
	client    *s3.Client
	presigner *s3.PresignClient
	bucket    string
}

func NewS3Storage(client *s3.Client, bucket string) *s3Storage {
	return &s3Storage{
		client:    client,
		presigner: s3.NewPresignClient(client),
		bucket:    bucket,
	}
}

func (s *s3Storage) GeneratePresignedGetURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	if key == "" {
		return "", errors.New("storage: key cannot be empty")
	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	req, err := s.presigner.PresignGetObject(ctx, params, func(po *s3.PresignOptions) {
		po.Expires = expires
	})

	if err != nil {
		return "", fmt.Errorf("presign get: %w", err)
	}

	return req.URL, nil
}

func (s *s3Storage) GeneratePresignedPutURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	if key == "" {
		return "", errors.New("storage: key cannot be empty")
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	req, err := s.presigner.PresignPutObject(ctx, params, func(po *s3.PresignOptions) {
		po.Expires = expires
	})

	if err != nil {
		return "", fmt.Errorf("presign put: %w", err)
	}

	return req.URL, nil
}
