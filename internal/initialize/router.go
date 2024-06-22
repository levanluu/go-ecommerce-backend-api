package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/levanluu/go-ecommerce-backend-api/internal/controller"
	"github.com/levanluu/go-ecommerce-backend-api/internal/middleware"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AAA")
		c.Next()
		fmt.Println("After --> AAA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> CC")
		c.Next()
		fmt.Println("After --> CC")
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthenMiddleware(), BB(), CC())
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/user", controller.NewUserController().GetUserByID)
	}

	return r
}
