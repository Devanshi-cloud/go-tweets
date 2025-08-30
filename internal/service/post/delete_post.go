package post

import (
	"context"
	"errors"
	
	"net/http"
	"time"
)

func (r *postService ) DeletePost(ctx context.Context, postID, userID int64) (int, error) {
	// check post was exists
	postExists, err := r.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExists == nil {
		return http.StatusNotFound, errors.New("tweet not found")
	}

	if postExists.UserID != userID {
		return http.StatusNotFound, errors.New("tweet not found")
	}
	// delete post
	err = r.postRepo.SoftDeletePost(ctx, postID, time.Now())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusOK, nil
}