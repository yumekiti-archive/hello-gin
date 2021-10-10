package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Api() {
	router:= gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":3000")
}