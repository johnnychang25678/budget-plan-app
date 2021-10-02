package controllers

import (
	"budget-plan-app/backend/models"
	"budget-plan-app/backend/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	repo *repositories.MemberRepo
}

func NewMemberController() *MemberController {
	repo := repositories.NewMemberRepo()
	return &MemberController{
		repo: repo,
	}
}

func (m *MemberController) CreateMember(c *gin.Context) {
	var member models.Member
	err := c.BindJSON(&member)
	if err != nil {
		fmt.Println("CreateMember controller err: ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = m.repo.Create(member)
	if err != nil {
		fmt.Println("member repo create error. email: ", member.Email)
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{"response": "ok"})
}
