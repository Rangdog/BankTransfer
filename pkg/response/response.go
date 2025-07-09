package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:   0,
		Status: "success",
		Data:   data,
	})
}

func Error(c *gin.Context, httpCode int, appCode int, message string) {
	c.JSON(httpCode, APIResponse{
		Code:    appCode,
		Status:  "error",
		Message: message,
	})
}
