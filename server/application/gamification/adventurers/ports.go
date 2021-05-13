package adventurers

import (
	"github.com/google/uuid"

	"tasquest.com/server/commons"
)

type AdventurerService interface {
	CreateAdventurer(command CreateAdventurer) (Adventurer, error)
	UpdateAdventurer(adventurerID uuid.UUID, command UpdateAdventurer) (Adventurer, error)
	UpdateCharacter(adventurerID uuid.UUID, character Character) (Adventurer, error)
}

type AdventurerPersistence interface {
	Save(adventurer Adventurer) (Adventurer, error)
	Update(adventurer Adventurer) (Adventurer, error)
}

type AdventurerFinder interface {
	FindByID(id uuid.UUID) (Adventurer, error)
	FindByUser(userID uuid.UUID) (Adventurer, error)
	FindOneByFilter(filter commons.SqlFilter) (Adventurer, error)
	FindAllByFilter(filter commons.SqlFilter) ([]Adventurer, error)
}
