package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoDB) Query(collectionName string, body []byte) ([]bson.M, error) {
	query, err := createQueryData(body)
	if err != nil {
		return nil, err
	}

	collection := m.Database.Collection(collectionName)

	findOptions := options.Find()

	if query.Sort != nil {
		findOptions.SetSort(query.Sort)
	}
	if query.Skip != 0 {
		findOptions.SetSkip(query.Skip)
	}
	if query.Limit != 0 {
		findOptions.SetLimit(query.Limit)
	}

	cursor, err := collection.Find(context.TODO(), query.Find, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var data []bson.M
	cursor.All(context.TODO(), &data)

	return data, nil
}
