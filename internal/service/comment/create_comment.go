package comment

import (
	"context"
	"errors"
	"fmt"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *commentService) CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID int64) (int, error) {
	// check tweet exist
	fmt.Printf("PostID: %d", req.PostID)
	postExist, err := s.postRepo.GetPostByID(ctx, req.PostID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExist == nil {
		return http.StatusNotFound, errors.New("tweet not found")
	}

	// store comment
	now := time.Now()
	err = s.commentRepo.StoreComment(ctx, &model.CommentModel{
		PostID:    req.PostID,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusCreated, nil
}
