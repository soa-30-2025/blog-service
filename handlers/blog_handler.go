package handlers

import (
    "context"
    pb "blog-service/proto/blog"
    "blog-service/services"
)

type BlogHandler struct {
    pb.UnimplementedBlogServiceServer
    Service *services.BlogService
}

func (h *BlogHandler) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
    err := h.Service.CreateBlog(req.Blog)
    if err != nil {
        return nil, err
    }
    return &pb.CreateBlogResponse{Blog: req.Blog}, nil
}

func (h *BlogHandler) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
    blog, err := h.Service.GetBlog(req.Id)
    if err != nil {
        return nil, err
    }
    return &pb.GetBlogResponse{Blog: blog}, nil
}
