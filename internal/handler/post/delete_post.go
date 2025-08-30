package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeletePost(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		userID = c.GetInt64("userID")
		PostIDStr = c.Param("post_id")
	)

	postID, err := strconv.ParseInt(PostIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post ID",
		})
		return
	}
	statusCode, err := h.postService.DeletePost(ctx, postID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"message": "tweet deleted successfully",
	})
}