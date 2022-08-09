package mongodb

import (
	"context"
)

func (m *MongoDB) Delete(collectionName string, params map[string]string) error {
	collection := m.Database.Collection(collectionName)
	filter := createFilter(params)

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
