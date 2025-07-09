package router

import (
	"BankTransfer/internal/handler"
	"BankTransfer/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	r.Use(middleware.RecoveryWithLogrus(logger))
	api := r.Group("/api")

	registerUserRoutes(api, userHandler)
}
