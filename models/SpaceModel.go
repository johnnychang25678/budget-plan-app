package models

type SpaceModel struct {
	MemberId   int    `json:"memberId" binding:"required"`
	SpaceTitle string `json:"spaceTitle" binding:"required"`
}
