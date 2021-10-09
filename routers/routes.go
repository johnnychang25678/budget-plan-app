package routers

import (
	"budget-plan-app/backend/controllers"

	"github.com/gin-gonic/gin"
)

var (
	memberController controllers.MemberController = *controllers.NewMemberController()
	spaceController  controllers.SpaceController  = *controllers.NewSpaceController()
)

func Routes(server *gin.Engine) {
	memberRoutes := server.Group("/member")
	{
		memberRoutes.POST("/signup", memberController.CreateMember)
	}

	spaceRoutes := server.Group("/spaces")
	{
		spaceRoutes.GET("/:memberId", spaceController.FindSpaces)
		spaceRoutes.POST("/", spaceController.CreateSpace)
		spaceRoutes.POST("/share", spaceController.ShareSpace)
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
