package controllers

import (
	"budget-plan-app/backend/models"
	"budget-plan-app/backend/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RepaymentPlanController struct {
	repo *repositories.RepaymentPlanRepo
}

func NewRepaymentPlanController() *RepaymentPlanController {
	repo := repositories.NewRepaymentRepo()
	return &RepaymentPlanController{
		repo: repo,
	}
}

func (r *RepaymentPlanController) CreateRepaymentPlan(c *gin.Context) {
	var repamymentPlan models.RepaymentPlan
	err := c.BindJSON(&repamymentPlan)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = r.repo.Create(repamymentPlan)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (r *RepaymentPlanController) GetRepaymentPlans(c *gin.Context) {
	spaceId, err := strconv.Atoi(c.Param("spaceId"))
	if err != nil {
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	plans, err := r.repo.GetPlansById(spaceId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, plans)
}
