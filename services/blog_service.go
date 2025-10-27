package services

import (
	models "blog-service/models"
	pb "blog-service/proto/blog"
	"blog-service/repository"
	"context"
)

type BlogService struct {
    Repo *repository.BlogRepository
}

func (s *BlogService) CreateBlog(ctx context.Context, blog *pb.Blog) (*models.Blog, error) {
    return s.Repo.Create(ctx, blog)
}

func (s *BlogService) GetBlog(ctx context.Context, id string) (*models.Blog, error) {
    return s.Repo.GetByID(ctx, id)
}

func (s *BlogService) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	return s.Repo.GetAllBlogs(ctx)
}


