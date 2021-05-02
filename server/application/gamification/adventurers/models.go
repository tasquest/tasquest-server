package adventurers

import (
	"time"

	"github.com/google/uuid"

	"tasquest.com/server/application/gamification/equipments"
)

type Adventurer struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	UserID    uuid.UUID `json:"userId" bson:"user_id"`
	Name      string    `json:"name" bson:"name"`
	Surname   string    `json:"surname" bson:"surname"`
	Birthday  time.Time `json:"birthday" bson:"birthday"`
	Character Character `json:"character" bson:"character"`
}

type Character struct {
	AdventurerID uuid.UUID          `json:"adventurerId" bson:"adventurer_id"`
	CharName     string             `json:"charName" bson:"char_name"`
	Title        string             `json:"title" bson:"title"`
	Level        int                `json:"level" bson:"level"`
	Experience   int64              `json:"leveling" bson:"leveling"`
	Equipment    CharacterEquipment `json:"equipments" bson:"equipments"`
	Guilds       []uuid.UUID        `json:"guilds" bson:"guilds"`
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
	Head      equipments.Equipment `json:"head" bson:"head"`
	Torso     equipments.Equipment `json:"torso" bson:"torso"`
	LeftHand  equipments.Equipment `json:"leftHand" bson:"left_hand"`
	RightRand equipments.Equipment `json:"rightHand" bson:"right_hand"`
	Waist     equipments.Equipment `json:"waist" bson:"waist"`
	Legs      equipments.Equipment `json:"legs" bson:"legs"`
	Feet      equipments.Equipment `json:"feet" bson:"feet"`
}
