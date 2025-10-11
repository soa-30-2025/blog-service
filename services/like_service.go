package services

import (
	"blog-service/repository"
	"context"
	"fmt"
)

type LikeService struct {
    Repo *repository.LikeRepository
}

func (s *LikeService) LikeBlog(ctx context.Context, blogID, userID string) error {
    return s.Repo.AddLike(ctx, blogID, userID)
}

func (s *LikeService) UnlikeBlog(ctx context.Context, blogID, userID string) error {
    return s.Repo.RemoveLike(ctx, blogID, userID)
}

func (s *LikeService) GetLikesCount(ctx context.Context, blogID string) (int, error) {
    return s.Repo.CountLikes(ctx, blogID)
}

func (s *LikeService) GetLikedUsers(ctx context.Context, blogID string) ([]string, error) {
    userIDs, err := s.Repo.GetLikedUserIDs(ctx, blogID)
    if err != nil {
        return nil, fmt.Errorf("failed to get liked users: %w", err)
    }
    return userIDs, nil
}

