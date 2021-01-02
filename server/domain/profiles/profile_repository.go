package profiles

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileRepository interface {
	Save(profile Profile) (Profile, error)
	FindByID(id string) (Profile, error)
	FindByUser(userID string) (Profile, error)
	FindByFilter(filter bson.M) (Profile, error)
}
