package repository

import (
	"context"
	"time"

	models "blog-service/models"
	pb "blog-service/proto/blog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CommentRepository struct {
    DB *pgxpool.Pool
}

func (r *CommentRepository) Create(ctx context.Context, comment *pb.Comment) (*models.Comment, error) {
    var c models.Comment
    query := `
        INSERT INTO comments (blog_id, user_id, text, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, blog_id, user_id, text, created_at, updated_at
    `
    row := r.DB.QueryRow(ctx, query,
        comment.BlogId,
        comment.UserId,
        comment.Text,
        time.Now(),
        time.Now(),
    )

    if err := row.Scan(&c.ID, &c.BlogID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt); err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *CommentRepository) GetByBlogID(ctx context.Context, blogID string) ([]models.Comment, error) {
    query := `SELECT id, blog_id, user_id, text, created_at, updated_at FROM comments WHERE blog_id = $1 ORDER BY created_at DESC`
    rows, err := r.DB.Query(ctx, query, blogID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []models.Comment
    for rows.Next() {
        var c models.Comment
        if err := rows.Scan(&c.ID, &c.BlogID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt); err != nil {
            return nil, err
        }
        comments = append(comments, c)
    }

    return comments, nil
}

func (r *CommentRepository) Update(ctx context.Context, id string, text string) (*models.Comment, error) {
	query := `
		UPDATE comments
		SET text = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, blog_id, user_id, text, created_at, updated_at
	`

	var c models.Comment
	row := r.DB.QueryRow(ctx, query, text, id)
	if err := row.Scan(&c.ID, &c.BlogID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return nil, err
	}

	return &c, nil
}

