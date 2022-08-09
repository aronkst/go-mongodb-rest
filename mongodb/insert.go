package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) Insert(collectionName string, body []byte) (bson.M, error) {
	collection := m.Database.Collection(collectionName)

	document, err := createDocument(body)
	if err != nil {
		return nil, err
	}

	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}

	_id := getIDResult(result)

	data, err := m.Show(collectionName, _id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
