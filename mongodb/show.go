package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MongoDB) Show(collectionName string, id string) (bson.M, error) {
	collection := m.Database.Collection(collectionName)

	objectID, err := createObjectID(id)
	if err != nil {
		return nil, err
	}

	var data bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&data)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return data, nil
}
