package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header["access_token"]
		fmt.Println(accessToken)
		c.Next()
	}
}
