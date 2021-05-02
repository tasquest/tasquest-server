package leveling

import "github.com/google/uuid"

const AdventurerLevelingTopic = ""

type AdventurerLevelingEvent struct {
	AdventurerID uuid.UUID `json:"adventurerId" binding:"required"`
	OldLevel     int       `json:"oldLevel" binding:"required"`
	NewLevel     int       `json:"newLevel" binding:"required"`
}
