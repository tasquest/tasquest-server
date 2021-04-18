package adventurers

import (
	"github.com/google/uuid"
	"time"
)

type CreateAdventurer struct {
	UserID   uuid.UUID `json:"userId" binding:"required"`
	Name     string    `json:"name" binding:"required,alpha"`
	Surname  string    `json:"surname" binding:"required,alpha"`
	Birthday time.Time `json:"birthday" binding:"required,datetime"`
}

type UpdateAdventurer struct {
	Name     string    `json:"name" binding:"required,alpha"`
	Surname  string    `json:"surname" binding:"required,alpha"`
	Birthday time.Time `json:"birthday" binding:"required,datetime"`
}

type UpdateExperience struct {
	AdventurerID uuid.UUID `json:"adventurerId" binding:"required"`
	Experience   int64     `json:"experience" binding:"required"`
}
