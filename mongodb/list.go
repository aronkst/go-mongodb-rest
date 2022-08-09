package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) List(collectionName string, params map[string]string) ([]bson.M, error) {
	collection := m.Database.Collection(collectionName)
	filter := createFilter(params)

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var data []bson.M
	cursor.All(context.TODO(), &data)

	return data, nil
}
