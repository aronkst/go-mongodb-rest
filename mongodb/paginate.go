package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoDB) Paginate(collectionName string, body []byte) (DataPaginate, error) {
	paginate, err := createQueryData(body)
	if err != nil {
		return DataPaginate{}, err
	}

	collection := m.Database.Collection(collectionName)

	findOptions := options.Find()
	skip := createPaginateSkip(paginate)

	findOptions.SetSkip(skip)
	findOptions.SetLimit(paginate.PerPage)
	if paginate.Sort != nil {
		findOptions.SetSort(paginate.Sort)
	}

	cursor, err := collection.Find(context.TODO(), paginate.Find, findOptions)
	if err != nil {
		return DataPaginate{}, err
	}

	defer cursor.Close(context.TODO())

	var data []bson.M
	cursor.All(context.TODO(), &data)

	count, err := collection.CountDocuments(context.TODO(), paginate.Find)
	if err != nil {
		return DataPaginate{}, err
	}

	totalPages := createPaginateTotalPages(paginate.PerPage, count)

	dataPaginate := DataPaginate{
		Page:       paginate.Page,
		TotalPages: totalPages,
		Data:       data,
	}

	return dataPaginate, nil
}
