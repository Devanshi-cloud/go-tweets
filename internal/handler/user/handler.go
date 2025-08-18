// Responsible for routing API requests to the appropriate handlers
// Handler methods are called directly from the Gin router, and they often contain a route list method to register all API routes

package user

import (
	"go-tweets/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api *gin.Engine
	validate *validator.Validate
	userService user.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) *Handler {
	return &Handler{
		api: api,
		validate: validate,
		userService: userService,
	}
}

func (h *Handler) RouteList() {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
	authRoute.POST("/login", h.Login)
}