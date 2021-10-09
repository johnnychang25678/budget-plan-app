package controllers

import (
	"budget-plan-app/backend/models"
	"budget-plan-app/backend/repositories"
	"budget-plan-app/backend/utils"
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
			"error": err.Error(),
		})
		return
	}

	err = s.repo.Create(space.MemberId, space.SpaceTitle, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"response": "ok"})
}

func (s *SpaceController) ShareSpace(c *gin.Context) {
	// client will bring in 1) member id and 2) space id he wants to share 3) the member id he wants to share with
	// check if this member has the space id && he's the owner
	type clientInputModel struct {
		MemberId         int `json:"memberId" binding:"required"`
		SpaceIdToShare   int `json:"spaceIdToShare" binding:"required"`
		MemeberIdToShare int `json:"memberIdToShare" binding:"required"`
	}
	var clientInput clientInputModel
	err := c.BindJSON(&clientInput)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	// if member does not own space, throw error
	spaces, err := s.repo.FindOwnedSpaces(clientInput.MemberId)
	fmt.Println(spaces)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	isMemberOwnSpace := utils.IsInSliceInt(spaces, clientInput.SpaceIdToShare)
	if !isMemberOwnSpace {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("memberId: %d does not own spaceId: %d", clientInput.MemberId, clientInput.SpaceIdToShare),
		})
		return
	}

	// if memberToShare already has this space shared, throw error
	memberToShareSpaces, err := s.repo.FindAll(clientInput.MemeberIdToShare)
	print("memberToShareSpaces:", memberToShareSpaces)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	isSpaceIdAlreadyShared := utils.IsInSliceInt(memberToShareSpaces, clientInput.SpaceIdToShare)
	if isSpaceIdAlreadyShared {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf(
				"memberId: %d already has spaceId: %d",
				clientInput.MemeberIdToShare,
				clientInput.SpaceIdToShare,
			),
		})
		return
	}
	// TODO: this portion will be changed to sending invitation email to the MemberIdToShare
	input := repositories.SpaceMemberInputModel{
		SpaceId:  clientInput.SpaceIdToShare,
		MemberId: clientInput.MemeberIdToShare,
		IsOwner:  false,
	}
	err = s.repo.AddToSpaceMember(input)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "ok",
	})
}

func (s *SpaceController) FindSpaces(c *gin.Context) {
	memberId, err := strconv.Atoi(c.Param("memberId"))
	fmt.Println("memberId: ", memberId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	spaces, err := s.repo.FindAll(memberId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"spaces": spaces,
	})
}
