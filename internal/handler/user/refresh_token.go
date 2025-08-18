package user
import (
	"go-tweets/internal/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

// creating handler function for user registration
func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Validate the request using the validator
	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID") // Assuming user ID is set in context after authentication
	token, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &req, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
	// Return the user ID and a success message
	// This will be returned in the response body as JSON
	// Example response: {"user_id": 123, "message": "User registered successfully"}
	// where 123 is the ID of the newly created user
	// You can customize the response structure as needed
	// Ensure to handle any potential errors and return appropriate HTTP status codes
}