package repository

import (
	"context"
	"errors"
	"social/internal/domain/post"

	"github.com/jackc/pgx/v5"
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

func (p *PostRepo) AddNewPost(ctx context.Context, postData *post.Post) (*post.Post, error) {

	createdPost := &post.Post{}

	query := `Insert Into "post" (user_id, title, description, tags) values ($1, $2, $3, $4 )
			returning user_id, title, description, tags`

	err := p.pool.QueryRow(ctx, query, postData.UserId, postData.Title, postData.Description, postData.Tags).
		Scan(&createdPost.UserId, &createdPost.Title, &createdPost.Description, &createdPost.Tags)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return createdPost, nil

}

func (p *PostRepo) GetPostById(ctx context.Context, id int64) (*post.Post, error) {
	post := &post.Post{}

	query := `Select id, user_id, title, description, tags, created_at from post where id=$1`
	err := p.pool.QueryRow(ctx, query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Description, &post.Tags, &post.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return post, err
}

func (p *PostRepo) GetPostsByUserId(ctx context.Context, UserId int64) ([]*post.Post, error) {

	query := `Select id, user_id, title, description, tags, created_at from post where user_id=$1`

	rows, err := p.pool.Query(ctx, query, UserId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*post.Post

	for rows.Next() {
		post := &post.Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Description, &post.Tags, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		err = rows.Err()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, err
}

func (p *PostRepo) UpdatePost() {

}
