package handler

import (
	"BankTransfer/internal/dto"
	userS "BankTransfer/internal/service/user"
	"BankTransfer/pkg/errorcode"
	"BankTransfer/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *userS.UserService
}

func MewUserHandler(service *userS.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.RegisterRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, errorcode.ErrInvalidParameter, "Invalid request body")
		return
	}
	user, err := h.service.RegisterUser(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, errorcode.ErrInternalServer, err.Error())
		return
	}
	response.Success(c, user)
}
