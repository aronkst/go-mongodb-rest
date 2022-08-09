package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MongoDB) Show(collectionName string, _id string) (bson.M, error) {
	collection := m.Database.Collection(collectionName)

	objectId, err := createObjectID(_id)
	if err != nil {
		return nil, err
	}

	var data bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}
