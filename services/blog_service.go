package services

import (
    pb "blog-service/proto/blog"
    "blog-service/repository"
)

type BlogService struct {
    Repo *repository.BlogRepository
}

func (s *BlogService) CreateBlog(blog *pb.Blog) error {
    // ovde bi mogla dodati validaciju, logovanje itd.
    return s.Repo.Create(blog)
}

func (s *BlogService) GetBlog(id string) (*pb.Blog, error) {
    return s.Repo.GetByID(id)
}
