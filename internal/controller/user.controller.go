package controller

import (
	"fmt"

	"example.com/ecommerce/internal/service"
	"example.com/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	fmt.Println("---> User controller")
	user := uc.userService.GetInfoUser()
	response.SuccessResponse(c, 2001, []string{user})
}
