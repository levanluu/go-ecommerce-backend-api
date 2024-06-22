package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/levanluu/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
