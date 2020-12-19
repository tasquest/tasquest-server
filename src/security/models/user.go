package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Surname   string             `bson:"surname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Active    bool               `bson:"active"`
	Providers []SocialProvider   `bson:"providers"`
}
