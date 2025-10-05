package repository

import (
    "database/sql"
    pb "blog-service/proto/blog"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type BlogRepository struct {
    DB *sql.DB
}

func (r *BlogRepository) Create(blog *pb.Blog) error {
    query := `INSERT INTO blogs (id, title, description, author_id, created_at) VALUES (?, ?, ?, ?, ?)`
    _, err := r.DB.Exec(query, blog.Id, blog.Title, blog.Description, blog.AuthorId, blog.CreatedAt.AsTime())
    return err
}

func (r *BlogRepository) GetByID(id string) (*pb.Blog, error) {
    query := `SELECT id, title, description, author_id, created_at FROM blogs WHERE id = ?`
    row := r.DB.QueryRow(query, id)

    var b pb.Blog
    var createdAt sql.NullTime
    err := row.Scan(&b.Id, &b.Title, &b.Description, &b.AuthorId, &createdAt)
    if err != nil {
        return nil, err
    }

    if createdAt.Valid {
        b.CreatedAt = timestamppb.New(createdAt.Time)
    }

    return &b, nil
}
