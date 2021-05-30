package fixtures

import (
	"time"

	"github.com/google/uuid"

	"tasquest.com/server/application/gamification/adventurers"
)

func NewAdventurer() adventurers.Adventurer {
	return adventurers.Adventurer{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		Name:      "John",
		Surname:   "Doe",
		Birthday:  time.Now(),
		Character: NewCharacter(),
	}
}

func NewCharacter() adventurers.Character {
	return adventurers.Character{
		ID:         uuid.New(),
		CharName:   "Joanerys",
		Title:      "The Unmatched",
		Level:      1,
		Experience: 10,
	}
}
