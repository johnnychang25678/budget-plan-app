package models

import "time"

type Member struct {
	Id        int
	Email     string `json:"email" binding:"required"`
	CreatedAt time.Time
}
