package services

import (
    pb "blog-service/proto/blog"
    models "blog-service/models"
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


