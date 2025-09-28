package storage

import (
	"context"
	"time"
)

type StorageService interface {
	GeneratePresignedPutURL(ctx context.Context, key string, expires time.Duration) (string, error)
	GeneratePresignedGetURL(ctx context.Context, key string, expires time.Duration) (string, error)
}
