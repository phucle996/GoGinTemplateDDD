package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 400 Bad Request
func RespondBadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Error:   err,
		Message: "bad request",
	})
}

// 401 Unauthorized
func RespondUnauthorized(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Error:   err,
		Message: "unauthorized",
	})
}

// 403 Forbidden
func RespondForbidden(c *gin.Context, err string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Error:   err,
		Message: "forbidden",
	})
}

// 404 Not Found
func RespondNotFound(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Error:   err,
		Message: "not found",
	})
}

// 500 Internal Server Error
func RespondInternalError(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Error:   err,
		Message: "internal server error",
	})
}
