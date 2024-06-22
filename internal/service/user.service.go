package service

import "github.com/levanluu/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInforUser() string {
	return us.userRepo.GetInforUser()
}
