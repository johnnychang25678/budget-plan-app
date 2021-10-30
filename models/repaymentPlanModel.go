package models

import "time"

type RepaymentPlan struct {
	SpaceId   int       `json:"spaceId" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	TotalCost int       `json:"totalCost" binding:"required"`
	DueDate   time.Time `json:"dueDate" binding:"required"`
	// YYYY-MM-DD HH:MM:SS

}
