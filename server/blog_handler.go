/*package server

import (
    "context"
    "log"
    "database/sql"
    "google.golang.org/protobuf/types/known/timestamppb"

    pb "blog-service/proto/blog"
)

type BlogServer struct {
    pb.UnimplementedBlogServiceServer
    DB *sql.DB
}

// Implementacija CreateBlog
func (s *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
    blog := req.Blog
    log.Println("Primljen request za kreiranje bloga:", blog.Title)

    query := `INSERT INTO blogs (id, title, description, author_id, created_at) VALUES (?, ?, ?, ?, ?)`
    _, err := s.DB.Exec(query, blog.Id, blog.Title, blog.Description, blog.AuthorId, blog.CreatedAt.AsTime())
    if err != nil {
        log.Println("Greška pri upisu u bazu:", err)
        return nil, err
    }

    return &pb.CreateBlogResponse{Blog: blog}, nil
}


// Implementacija GetBlog
func (s *BlogServer) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
    log.Println("Primljen request za dobijanje bloga ID:", req.Id)

    query := `SELECT id, title, description, author_id, created_at FROM blogs WHERE id = ?`
    row := s.DB.QueryRow(query, req.Id)

    var id, title, description, authorId string
    var createdAt sql.NullTime

    err := row.Scan(&id, &title, &description, &authorId, &createdAt)
    if err != nil {
        log.Println("Greška pri dohvatanju bloga:", err)
        return nil, err
    }

    blog := &pb.Blog{
        Id:          id,
        Title:       title,
        Description: description,
        AuthorId:    authorId,
        CreatedAt:   timestamppb.New(createdAt.Time),
    }

    return &pb.GetBlogResponse{Blog: blog}, nil
}

*/