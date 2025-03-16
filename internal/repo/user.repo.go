package repo

import "fmt"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	fmt.Println("---> User repo")
	return "LuuLe"
}
