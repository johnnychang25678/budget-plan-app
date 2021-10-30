package routers

import (
	"budget-plan-app/backend/controllers"

	"github.com/gin-gonic/gin"
)

var (
	memberController         controllers.MemberController
	spaceController          controllers.SpaceController
	repamymentPlanController controllers.RepaymentPlanController
	authController           controllers.AuthController
)

func initControllers() {
	memberController = *controllers.NewMemberController()
	spaceController = *controllers.NewSpaceController()
	repamymentPlanController = *controllers.NewRepaymentPlanController()
	authController = *controllers.NewAuthController()
}

func Routes(server *gin.Engine) {
	initControllers()
	authRoutes := server.Group("/auth")
	{
		authRoutes.GET("/", authController.HandleAuth)
		authRoutes.GET("/callback", authController.HandleCallback)
	}

	memberRoutes := server.Group("/member")
	{
		memberRoutes.POST("/signup", memberController.CreateMember)
	}

	// testRoutes := server.Group("/test")
	// {
	// 	testRoutes.Use(middlewares.Auth())
	// 	testRoutes.GET("/", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"status": "You're in!",
	// 		})
	// 	})
	// }

	spaceRoutes := server.Group("/spaces")
	{
		spaceRoutes.GET("/:memberId", spaceController.FindSpaces)
		spaceRoutes.POST("/", spaceController.CreateSpace)
		spaceRoutes.POST("/share", spaceController.ShareSpace)
	}
	planRoutes := server.Group("/repayment-plans")
	{
		planRoutes.POST("/", repamymentPlanController.CreateRepaymentPlan)
		planRoutes.GET("/:spaceId", repamymentPlanController.GetRepaymentPlans)
	}
	budgetRoutes := server.Group("/budgets")
	{
		budgetRoutes.GET("/", getBudgets)
		budgetRoutes.GET("/:budgetId", getBudgets)
	}
}
