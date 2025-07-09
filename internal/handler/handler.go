package handler

import (
	user "BankTransfer/internal/handler/user"
	"BankTransfer/internal/service"
)

type Handler struct {
	UserHandler *user.UserHandler
}

func NewHandler(service service.Services) *Handler {
	return &Handler{
		UserHandler: user.MewUserHandler(service.UserService),
	}
}
