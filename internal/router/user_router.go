package router

import (
	"BankTransfer/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(api *gin.RouterGroup, h *handler.UserHandler) {
	api.POST("/auth/register", h.Register)
}
