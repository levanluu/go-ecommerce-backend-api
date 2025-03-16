package routers

import (
	"fmt"

	c "example.com/ecommerce/internal/controller"
	"example.com/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -->AA")
		c.Next()
		fmt.Println("After -->AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -->BB")
		c.Next()
		fmt.Println("After -->BB")
	}
}

func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -->CC")
		c.Next()
		fmt.Println("After -->CC")
	}
}

func NewServer() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), BB(), CC())
	v1 := r.Group("/v1")
	v1.GET("/ping", c.NewPongController().Pong)
	v1.GET("/user", c.NewUserController().GetUserByID)
	return r
}
