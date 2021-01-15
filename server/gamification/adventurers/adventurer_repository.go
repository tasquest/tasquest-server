package adventurers

import (
	"go.mongodb.org/mongo-driver/bson"
)

type AdventurerRepository interface {
	Save(profile Adventurer) (Adventurer, error)
	Update(adventurer Adventurer) (Adventurer, error)
	FindByID(id string) (Adventurer, error)
	FindByUser(userID string) (Adventurer, error)
	FindByFilter(filter bson.M) (Adventurer, error)
}
