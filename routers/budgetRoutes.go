package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBudgets(c *gin.Context) {
	c.JSON(http.StatusOK, "get all budgets")
}

func getBudget(c *gin.Context) {
	c.JSON(http.StatusOK, "get budget detail")
}
