package repository

import (
	"context"
	"errors"
	"social/internal/domain/post"

	"github.com/Masterminds/squirrel"
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

func (p *PostRepo) AddNewPost(ctx context.Context, postData *post.Post) error {

	query := `Insert Into "post" (user_id, title, description, tags, media_urls) values ($1, $2, $3, $4, $5)
			returning user_id, title, description, tags`

	_, err := p.pool.Exec(ctx, query, postData.UserId, postData.Title, postData.Description, postData.Tags, postData.MediaURLs)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil

}

func (p *PostRepo) GetPostById(ctx context.Context, id int64) (*post.Post, error) {
	post := &post.Post{}

	query := `Select id, user_id, title, description, tags, media_urls, created_at from post where id=$1`
	err := p.pool.QueryRow(ctx, query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Description, &post.Tags, &post.MediaURLs, &post.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return post, err
}

func (p *PostRepo) GetPostsByUserId(ctx context.Context, UserId int64, filter post.PostByUserIdFilter) ([]*post.Post, error) {
	qb := squirrel.Select("id", "user_id", "title", "description", "tags", "media_urls", "created_at").
		From("post").
		Where(squirrel.Eq{"user_id": UserId}).
		PlaceholderFormat(squirrel.Dollar)

	// Title filter
	if filter.Filters.Title != "" {
		qb = qb.Where(squirrel.ILike{"title": "%" + filter.Filters.Title + "%"})
	}

	// Desc filter
	if filter.Filters.Desc != "" {
		qb = qb.Where(squirrel.ILike{"description": "%" + filter.Filters.Desc + "%"})
	}

	// Search across title OR desc
	if filter.Search != "" {
		qb = qb.Where(squirrel.Or{
			squirrel.ILike{"title": "%" + filter.Search + "%"},
			squirrel.ILike{"description": "%" + filter.Search + "%"},
		})
	}

	// Tags filter
	if len(filter.Filters.Tags) > 0 {
		qb = qb.Where("tags @> ?", filter.Filters.Tags)
	}

	// Sorting
	qb = qb.OrderBy("created_at " + filter.Sort)

	// Pagination
	if filter.Limit <= 0 {
		filter.Limit = 10 // Default limit
	}

	if filter.Page < 1 {
		filter.Page = 1 // Default page
	}

	qb = qb.Limit(uint64(filter.Limit)).Offset(uint64((filter.Page - 1) * filter.Limit))

	// Build the query
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.pool.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*post.Post

	for rows.Next() {
		post := &post.Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Description, &post.Tags, &post.MediaURLs, &post.CreatedAt)
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
