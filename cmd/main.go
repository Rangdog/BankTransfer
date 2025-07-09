package cmd

import (
	"BankTransfer/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	router.SetupRoutes(r)
}
