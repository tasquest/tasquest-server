package utils

import (
	"context"
	"fmt"
	"github.com/orlangure/gnomock"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo(container *gnomock.Container) *mongo.Client {
	addr := container.DefaultAddress()
	uri := fmt.Sprintf("mongodb://%s:%s@%s", "tasquest", "tasquest", addr)
	clientOptions := mongoOptions.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		panic(err)
	}

	return client
}
