package routers

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	spaceRoutes := server.Group("/space")
	{
		spaceRoutes.GET("/", getSpace)
	}
	planRoutes := server.Group("/plans")
	{
		planRoutes.GET("/", getPlans)
		planRoutes.GET("/:planId", getPlan)
	}
	budgetRoutes := server.Group("/budgets")
	{
		budgetRoutes.GET("/", getBudgets)
		budgetRoutes.GET("/:budgetId", getBudgets)
	}
}
