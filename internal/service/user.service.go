package service

import (
	"fmt"

	"example.com/ecommerce/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	fmt.Println("---> User servicer")
	return us.userRepo.GetInfoUser()
}
