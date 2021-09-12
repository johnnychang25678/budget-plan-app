package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSpace(c *gin.Context) {
	c.JSON(http.StatusOK, "Information about your space")
}
