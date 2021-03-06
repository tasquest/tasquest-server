package adventurers

import (
	"github.com/google/uuid"
	"tasquest.com/server/commons"
)

type AdventurerService interface {
	CreateAdventurer(command CreateAdventurer) (Adventurer, error)
	UpdateAdventurer(adventurerID uuid.UUID, command UpdateAdventurer) (Adventurer, error)
	UpdateExperience(command UpdateExperience) (Adventurer, error)
}

type AdventurerPersistence interface {
	Save(adventurer Adventurer) (Adventurer, error)
	Update(adventurer Adventurer) (Adventurer, error)
}

type AdventurerFinder interface {
	FindByID(id uuid.UUID) (Adventurer, error)
	FindByUser(userID uuid.UUID) (Adventurer, error)
	FindByFilter(filter commons.Map) (Adventurer, error)
}
