package handlers

import (
	"context"
	pb "blog-service/proto/blog"
	"blog-service/services"
    "google.golang.org/protobuf/types/known/timestamppb"
    "time"

)

type BlogHandler struct {
	pb.UnimplementedBlogServiceServer
	Service *services.BlogService
}

func (h *BlogHandler) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := &pb.Blog{
		Title:       req.Title,
		Description: req.Description,
		AuthorId:    req.AuthorId,
		CreatedAt:   timestamppb.New(time.Now()),
	}

	createdBlog, err := h.Service.CreateBlog(ctx, blog)
	if err != nil {
		return nil, err
	}

	// Mapiraj createdBlog u pb.Blog za response
	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:          createdBlog.ID,
			Title:       createdBlog.Title,
			Description: createdBlog.Description,
			AuthorId:    createdBlog.AuthorId,
			CreatedAt:   timestamppb.New(createdBlog.CreatedAt),
		},
	}, nil
}

func (h *BlogHandler) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	blog, err := h.Service.GetBlog(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetBlogResponse{
		Blog: &pb.Blog{
			Id:          blog.ID,
			Title:       blog.Title,
			Description: blog.Description,
			AuthorId:    blog.AuthorId,
			CreatedAt:   timestamppb.New(blog.CreatedAt),
		},
	}, nil
}
