package guilds

import (
	"github.com/google/uuid"
)

type Guild struct {
	ID      uuid.UUID   `json:"id" bson:"_id"`
	Name    string      `json:"name" bson:"name"`
	Motto   string      `json:"motto" bson:"motto"`
	Symbol  string      `json:"symbol" bson:"symbol"`
	Members []uuid.UUID `json:"members" bson:"members"`
	Ranks   []Rank      `json:"ranks" bson:"ranks"`
}

type Rank struct {
}
