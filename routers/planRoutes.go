package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPlans(c *gin.Context) {
	c.JSON(http.StatusOK, "All the plans within the space")
}

func getPlan(c *gin.Context) {
	c.JSON(http.StatusOK, "Detail of a single plan")
}
