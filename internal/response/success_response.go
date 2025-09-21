package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 200 OK
func RespondSuccess(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Message: message,
		Data:    data,
	})
}

// 201 Created
func RespondCreated(c *gin.Context, data any, message string) {
	c.JSON(http.StatusCreated, APIResponse{
		Message: message,
		Data:    data,
	})
}
