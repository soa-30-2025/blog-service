package services

import (
	models "blog-service/models"
	pb "blog-service/proto/blog"
	"blog-service/repository"
	"context"
)

type CommentService struct {
    Repo *repository.CommentRepository
}

func (s *CommentService) CreateComment(ctx context.Context, comment *pb.Comment) (*models.Comment, error) {
    return s.Repo.Create(ctx, comment)
}

func (s *CommentService) GetCommentsByBlog(ctx context.Context, blogID string) ([]models.Comment, error) {
    return s.Repo.GetByBlogID(ctx, blogID)
}
