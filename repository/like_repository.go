package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LikeRepository struct {
    DB *pgxpool.Pool
}

func (r *LikeRepository) AddLike(ctx context.Context, blogID, userID string) error {
    query := `INSERT INTO likes (blog_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
    _, err := r.DB.Exec(ctx, query, blogID, userID)
    return err
}

func (r *LikeRepository) RemoveLike(ctx context.Context, blogID, userID string) error {
    query := `DELETE FROM likes WHERE blog_id = $1 AND user_id = $2`
    _, err := r.DB.Exec(ctx, query, blogID, userID)
    return err
}

func (r *LikeRepository) CountLikes(ctx context.Context, blogID string) (int, error) {
    query := `SELECT COUNT(*) FROM likes WHERE blog_id = $1`
    var count int
    err := r.DB.QueryRow(ctx, query, blogID).Scan(&count)
    return count, err
}
