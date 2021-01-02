package profiles

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tasquest.com/server/domain/gamification/characters"
)

type Profile struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	UserID    primitive.ObjectID   `json:"userId" bson:"user_id"`
	Name      string               `json:"name" bson:"name"`
	Surname   string               `json:"surname" bson:"surname"`
	Birthday  primitive.DateTime   `json:"birthday" bson:"birthday"`
	Character characters.Character `json:"character" bson:"character"`
}
