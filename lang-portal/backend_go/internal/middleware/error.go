package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandler middleware handles errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Handle different types of errors
			switch err.Type {
			case gin.ErrorTypeBind:
				c.JSON(http.StatusBadRequest, ErrorResponse{
					Error: "Invalid request parameters",
				})
			default:
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error: "Internal server error",
				})
			}

			// Abort the request
			c.Abort()
		}
	}
}

// NotFound handles 404 errors
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error: "Resource not found",
		})
	}
}
