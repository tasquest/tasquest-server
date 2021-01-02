package characters

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasquest.com/server/domain/gamification/equipments"
)

type Character struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	CharName   string               `json:"charName" bson:"char_name"`
	Title      string               `json:"title" bson:"title"`
	Level      int                  `json:"level" bson:"level"`
	Experience int                  `json:"experience" bson:"experience"`
	Equipment  CharacterEquipment   `json:"equipments" bson:"equipments"`
	Guilds     []primitive.ObjectID `json:"guilds" bson:"guilds"`
}

type CharacterEquipment struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	Head      equipments.Equipment `json:"head" bson:"head"`
	Torso     equipments.Equipment `json:"torso" bson:"torso"`
	LeftHand  equipments.Equipment `json:"leftHand" bson:"left_hand"`
	RightRand equipments.Equipment `json:"rightHand" bson:"right_hand"`
	Waist     equipments.Equipment `json:"waist" bson:"waist"`
	Legs      equipments.Equipment `json:"legs" bson:"legs"`
	Feet      equipments.Equipment `json:"feet" bson:"feet"`
}
