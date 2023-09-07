package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/mock-api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Mock API response",
		})
	})

	router.Run(":1000")
}
