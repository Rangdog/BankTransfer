package middleware

import (
	"BankTransfer/pkg/errorcode"
	"BankTransfer/pkg/response"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryWithLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithFields(logrus.Fields{
					"panic": err,
					"stack": string(debug.Stack()),
				}).Error("PANIC RECOVERED")

				response.Error(c, http.StatusInternalServerError, errorcode.ErrInternalServer, errorToString(err))
				c.Abort() // Quan trọng để ngăn handler tiếp tục sau panic
			}
		}()
		c.Next()
	}
}

func errorToString(err interface{}) string {
	if e, ok := err.(string); ok {
		return e
	}
	return "internal server error"
}
