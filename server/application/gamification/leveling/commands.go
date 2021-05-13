package leveling

import "github.com/google/uuid"

type AwardExperience struct {
	AdventurerID uuid.UUID `json:"adventurerId" binding:"required"`
	Experience   int64     `json:"leveling" binding:"required"`
}

type CreateLevel struct {
	NewTopExp int64 `json:"toExp" bson:"toExp"`
}
