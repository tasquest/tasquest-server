package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SocialProvider struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Token string             `bson:"token"`
}
