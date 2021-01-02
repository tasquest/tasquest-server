package guilds

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guild struct {
	ID      primitive.ObjectID   `json:"id" bson:"_id"`
	Name    string               `json:"name" bson:"name"`
	Motto   string               `json:"motto" bson:"motto"`
	Symbol  string               `json:"symbol" bson:"symbol"`
	Members []primitive.ObjectID `json:"members" bson:"members"`
	Ranks   []Rank               `json:"ranks" bson:"ranks"`
}

type Rank struct {
}
