package adventurers

import (
	"time"

	"github.com/google/uuid"
)

type Adventurer struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthday  time.Time `json:"birthday"`
	Character Character `json:"character" gorm:"foreignKey:AdventurerID"`
}

type Character struct {
	ID           uuid.UUID `json:"id"`
	CharName     string    `json:"charName"`
	Title        string    `json:"title"`
	Level        int       `json:"level"`
	Experience   int64     `json:"leveling"`
	AdventurerID uuid.UUID `json:"adventurerId"`
	// Equipment  CharacterEquipment `json:"equipments" bson:"equipments"`
	// Guilds     []uuid.UUID        `json:"guilds" bson:"guilds"`
}

func (char Character) NewLevel(level int) Character {
	char.Level = level
	return char
}

func (char Character) IncreaseExperience(experience int64) Character {
	char.Experience += experience
	return char
}

func (char Character) DecreaseExperience(experience int64) Character {
	char.Experience -= experience
	return char
}

type CharacterEquipment struct {
	Head      uuid.UUID `json:"head" bson:"head"`
	Torso     uuid.UUID `json:"torso" bson:"torso"`
	LeftHand  uuid.UUID `json:"leftHand" bson:"left_hand"`
	RightRand uuid.UUID `json:"rightHand" bson:"right_hand"`
	Waist     uuid.UUID `json:"waist" bson:"waist"`
	Legs      uuid.UUID `json:"legs" bson:"legs"`
	Feet      uuid.UUID `json:"feet" bson:"feet"`
}
