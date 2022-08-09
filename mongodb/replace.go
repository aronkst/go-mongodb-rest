package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) Replace(collectionName string, params map[string]string, body []byte) ([]bson.M, error) {
	collection := m.Database.Collection(collectionName)
	filter := createFilter(params)

	document, err := createDocument(body)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var dataReplace []bson.M
	cursor.All(context.TODO(), &dataReplace)

	for _, dR := range dataReplace {
		_, err = collection.ReplaceOne(context.TODO(), bson.M{"_id": dR["_id"]}, document)
		if err != nil {
			return nil, err
		}
	}

	cursor, err = collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var data []bson.M
	cursor.All(context.TODO(), &data)

	return data, nil
}
