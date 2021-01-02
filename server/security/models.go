package security

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Active    bool               `json:"active" bson:"active"`
	Providers []SocialProvider   `json:"providers" bson:"providers"`
}

type SocialProvider struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Token string             `json:"token" bson:"token"`
}
