package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: complete auth middleware
		// accessToken := c.Request.Header["access_token"]
		// refreshToken = c.Request.Header["refresh_token"]

		// fmt.Println(accessToken)
		c.Next()
	}
}
