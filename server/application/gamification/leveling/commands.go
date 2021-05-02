package leveling

import "github.com/google/uuid"

type AwardExperience struct {
	AdventurerID uuid.UUID `json:"adventurerId" binding:"required"`
	Experience   int64     `json:"leveling" binding:"required"`
}

type CreateLevel struct {
	Level   int   `json:"level" bson:"level"`
	FromExp int64 `json:"fromExp" bson:"fromExp"`
	ToExp   int64 `json:"toExp" bson:"toExp"`
}

type DeleteLevel struct {
	Level int `json:"level" bson:"level"`
}
