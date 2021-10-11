package routes

import (
    "github.com/gin-gonic/gin"

	"api/todo"
)

func Api() {
	router:= gin.Default()
	router.GET("/", todo.Index())
	router.Run(":3000")
}