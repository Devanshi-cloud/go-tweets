package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-tweets/internal/service/post"
)

type Handler struct {
	api *gin.Engine
	validate *validator.Validate
	postService post.PostService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService) *Handler {
	return &Handler{
		api: api,
		validate: validate,
		postService: postService,
	}
}

func (h * Handler) RouteList(secretKey string) {
	
}