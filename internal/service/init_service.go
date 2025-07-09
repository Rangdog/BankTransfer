package service

import (
	"BankTransfer/internal/repository"
	userS "BankTransfer/internal/service/user"
)

type Services struct {
	UserService *userS.UserService
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		UserService: userS.NewUserService(repos.UserRepository),
	}
}
