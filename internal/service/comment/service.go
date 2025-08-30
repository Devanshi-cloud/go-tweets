package comment

import (
	"context"
	"go-tweets/internal/config"
	"go-tweets/internal/dto"
	"go-tweets/internal/repository/comment"
	"go-tweets/internal/repository/post"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID int64) (int, error)
	UpdateComment(ctx context.Context, req *dto.StoreCommentRequest, commentID, userID int64) (int, error)
	DeleteComment(ctx context.Context, commentID, userID int64) (int, error)
}

type commentService struct {
	cfg         *config.Config
	commentRepo comment.CommentRepository
	postRepo    post.PostRepository
}

func NewCommentService(cfg *config.Config, commentRepo comment.CommentRepository, postRepo post.PostRepository) CommentService {
	return &commentService{
		cfg:         cfg,
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}
