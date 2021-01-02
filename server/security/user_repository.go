package security

import "go.mongodb.org/mongo-driver/bson"

type UserRepository interface {
	Save(user User) (User, error)
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	FindByFilter(filter bson.M) (User, error)
}
