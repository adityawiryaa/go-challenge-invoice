package main

import (
	"go-invoice/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	router.GET("/invoice", handler.GetInvoice)
	router.Run()
}
