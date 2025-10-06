package handlers

import (
	"context"
	"time"

	pb "blog-service/proto/blog"
	"blog-service/services"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BlogHandler struct {
	pb.UnimplementedBlogServiceServer
	BlogService    *services.BlogService
	CommentService *services.CommentService
}

func (h *BlogHandler) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := &pb.Blog{
		Title:       req.Title,
		Description: req.Description,
		AuthorId:    req.AuthorId,
		CreatedAt:   timestamppb.New(time.Now()),
	}

	createdBlog, err := h.BlogService.CreateBlog(ctx, blog)
	if err != nil {
		return nil, err
	}

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
	blog, err := h.BlogService.GetBlog(ctx, req.Id)
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

func (h *BlogHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	comment := &pb.Comment{
		BlogId:    req.BlogId,
		UserId:    req.UserId,
		Text:      req.Text,
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}

	createdComment, err := h.CommentService.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{
		Comment: &pb.Comment{
			Id:        createdComment.ID,
			BlogId:    createdComment.BlogID,
			UserId:    createdComment.UserID,
			Text:      createdComment.Text,
			CreatedAt: timestamppb.New(createdComment.CreatedAt),
			UpdatedAt: timestamppb.New(createdComment.UpdatedAt),
		},
	}, nil
}

func (h *BlogHandler) GetCommentsByBlog(ctx context.Context, req *pb.GetCommentsByBlogRequest) (*pb.GetCommentsByBlogResponse, error) {
	comments, err := h.CommentService.GetCommentsByBlog(ctx, req.BlogId)
	if err != nil {
		return nil, err
	}

	var pbComments []*pb.Comment
	for _, c := range comments {
		pbComments = append(pbComments, &pb.Comment{
			Id:        c.ID,
			BlogId:    c.BlogID,
			UserId:    c.UserID,
			Text:      c.Text,
			CreatedAt: timestamppb.New(c.CreatedAt),
			UpdatedAt: timestamppb.New(c.UpdatedAt),
		})
	}

	return &pb.GetCommentsByBlogResponse{
		Comments: pbComments,
	}, nil
}

func (h *BlogHandler) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	updatedComment, err := h.CommentService.UpdateComment(ctx, req.Id, req.Text)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCommentResponse{
		Comment: &pb.Comment{
			Id:        updatedComment.ID,
			BlogId:    updatedComment.BlogID,
			UserId:    updatedComment.UserID,
			Text:      updatedComment.Text,
			CreatedAt: timestamppb.New(updatedComment.CreatedAt),
			UpdatedAt: timestamppb.New(updatedComment.UpdatedAt),
		},
	}, nil
}

