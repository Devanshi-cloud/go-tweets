package post

import (
	"go-tweets/internal/config"
	"go-tweets/internal/repository/post"
)

type PostService interface {

}

type postService struct {
	cfg *config.Config
	postRepo post.PostRepository
}

func NewPostService(cfg *config.Config, postRepo post.PostRepository) PostService {
	return &postService{
		cfg:      cfg,
		postRepo: postRepo,
	}
}