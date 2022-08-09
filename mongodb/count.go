package mongodb

import (
	"context"
)

func (m *MongoDB) Count(collectionName string, body []byte) (DataCount, error) {
	collection := m.Database.Collection(collectionName)

	query, err := createQueryData(body)
	if err != nil {
		return DataCount{}, err
	}

	count, err := collection.CountDocuments(context.TODO(), query.Find)
	if err != nil {
		return DataCount{}, err
	}

	data := DataCount{Count: count}

	return data, nil
}
