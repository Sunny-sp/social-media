package repository

import (
	"social/internal/domain/post"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepo struct {
	pool *pgxpool.Pool
}

func NewPostRepo(pool *pgxpool.Pool) post.PostRepository {
	return &PostRepo{
		pool: pool,
	}
}

func (p *PostRepo) AddNewPost() {

}

func (p *PostRepo) UpdatePost() {

}
