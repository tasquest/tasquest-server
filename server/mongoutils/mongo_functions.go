package mongoutils

import (
	"context"
	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Save(collection *mongo.Collection, entity interface{}, bindTo interface{}) (interface{}, error) {
	insertResult, err := collection.InsertOne(context.Background(), entity)

	if err != nil {
		return bindTo, errors.WithStack(err)
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).String()

	return FindByID(collection, insertedID, bindTo)
}

func Update(collection *mongo.Collection, id primitive.ObjectID, entity interface{}, bindTo interface{}) (interface{}, error) {
	insertResult, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, entity)

	if err != nil {
		return bindTo, errors.WithStack(err)
	}

	insertedID := insertResult.UpsertedID.(primitive.ObjectID).String()

	return FindByID(collection, insertedID, bindTo)
}

func Delete(collection *mongo.Collection, entity interface{}) (int64, error) {
	deleteResult, err := collection.DeleteOne(context.Background(), entity)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return deleteResult.DeletedCount, nil
}

func DeleteByID(collection *mongo.Collection, id string, bindTo interface{}) (interface{}, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	deleteResult := collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID})
	err := deleteResult.Decode(&bindTo)
	return bindTo, err
}

func FindByID(collection *mongo.Collection, id string, bindTo interface{}) (interface{}, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return FindByFilter(collection, bson.M{"_id": objectID}, bindTo)
}

func FindByFilter(collection *mongo.Collection, filter bson.M, bindTo interface{}) (interface{}, error) {
	result := collection.FindOne(context.Background(), filter)
	err := result.Decode(&bindTo)

	if err != nil {
		return bindTo, errors.WithStack(err)
	}

	return bindTo, nil
}
