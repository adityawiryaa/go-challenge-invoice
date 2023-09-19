package main

import (
	"fmt"
	"go-invoice/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/invoice", handler.GetInvoice)

	port := "6005"
	fmt.Printf("Server is running on port %s...\n", port)
	router.Run(":" + port)
}
