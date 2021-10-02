package controllers

import (
	"budget-plan-app/backend/models"
	"budget-plan-app/backend/repositories"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SpaceController struct {
	repo *repositories.SpaceRepo
}

func NewSpaceController() *SpaceController {
	repo := repositories.NewSpaceRepo()
	return &SpaceController{
		repo: repo,
	}
}

func (s *SpaceController) CreateSpace(c *gin.Context) {
	var space models.SpaceModel
	err := c.BindJSON(&space)
	if err != nil {
		fmt.Println("CreateSpace controller err: ", err)
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	err = s.repo.Create(space.MemberId, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"response": "ok"})
}

func (s *SpaceController) FindSpaces(c *gin.Context) {
	memberId, err := strconv.Atoi(c.Param("memberId"))
	fmt.Println("memberId: ", memberId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	spaces, err := s.repo.FindAll(memberId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"spaces": spaces,
	})
}
