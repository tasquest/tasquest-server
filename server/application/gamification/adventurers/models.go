package adventurers

import (
	"github.com/google/uuid"
	"tasquest.com/server/application/gamification/equipments"
	"time"
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
	ID         uuid.UUID          `json:"id" bson:"_id"`
	CharName   string             `json:"charName" bson:"char_name"`
	Title      string             `json:"title" bson:"title"`
	Level      int                `json:"level" bson:"level"`
	Experience int64              `json:"experience" bson:"experience"`
	Equipment  CharacterEquipment `json:"equipments" bson:"equipments"`
	Guilds     []uuid.UUID        `json:"guilds" bson:"guilds"`
}

type CharacterEquipment struct {
	ID        uuid.UUID            `json:"id" bson:"_id"`
	Head      equipments.Equipment `json:"head" bson:"head"`
	Torso     equipments.Equipment `json:"torso" bson:"torso"`
	LeftHand  equipments.Equipment `json:"leftHand" bson:"left_hand"`
	RightRand equipments.Equipment `json:"rightHand" bson:"right_hand"`
	Waist     equipments.Equipment `json:"waist" bson:"waist"`
	Legs      equipments.Equipment `json:"legs" bson:"legs"`
	Feet      equipments.Equipment `json:"feet" bson:"feet"`
}
