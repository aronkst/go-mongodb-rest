package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) Update(collectionName string, params map[string]string, body []byte) ([]bson.M, error) {
	collection := m.Database.Collection(collectionName)
	filter := createFilter(params)

	document, err := createDocument(body)
	if err != nil {
		return nil, err
	}

	set := bson.M{"$set": document}

	_, err = collection.UpdateMany(context.TODO(), filter, set)
	if err != nil {
		return nil, err
	}

	data, err := m.List(collectionName, params)
	if err != nil {
		return nil, err
	}

	return data, nil
}
