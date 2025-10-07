package repository

import (
    models "blog-service/models"
    pb "blog-service/proto/blog"
    "time"
    "github.com/jackc/pgx/v5/pgxpool"
    "context"
)

type BlogRepository struct {
    DB *pgxpool.Pool
}

func (r *BlogRepository) Create(ctx context.Context, blog *pb.Blog) (*models.Blog, error) {
    var b models.Blog
	query := `
        INSERT INTO blogs (title, description, author_id, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, title, description, author_id, created_at
    `
	row := r.DB.QueryRow(ctx, query, blog.Title, blog.Description, blog.AuthorId, time.Now())
	
    if err := row.Scan(&b.ID, &b.Title, &b.Description, &b.AuthorId, &b.CreatedAt); err != nil {
        return nil, err
    }
    return &b, nil
}

func (r *BlogRepository) GetByID(ctx context.Context, id string) (*models.Blog, error) {
	query := `SELECT id, title, description, author_id, created_at FROM blogs WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, id)

	var b models.Blog
	err := row.Scan(&b.ID, &b.Title, &b.Description, &b.AuthorId, &b.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
