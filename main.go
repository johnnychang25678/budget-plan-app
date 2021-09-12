package main

import (
	"budget-plan-app/backend/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})

	routers.Routes(server)

	server.Run(":8080")
}
