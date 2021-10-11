package todo

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	}
}
