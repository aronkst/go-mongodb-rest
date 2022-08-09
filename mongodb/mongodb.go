package mongodb

import (
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Database *mongo.Database
}

func New(database *mongo.Database) *MongoDB {
	return &MongoDB{Database: database}
}

type DataPaginate struct {
	Page       int64    `json:"page"`
	TotalPages int64    `json:"total_pages"`
	Data       []bson.M `json:"data"`
}

type DataCount struct {
	Count int64 `json:"count"`
}

type queryData struct {
	Find    interface{} `json:"find"`
	Sort    interface{} `json:"sort"`
	Skip    int64       `json:"skip"`
	Limit   int64       `json:"limit"`
	Page    int64       `json:"page"`
	PerPage int64       `json:"perpage"`
}

func createFilter(params map[string]string) bson.M {
	filter := bson.M{}

	if len(params) > 0 {
		andCondition := bson.A{}

		for key, value := range params {
			expr := bson.M{"$expr": bson.M{"$eq": bson.A{bson.M{"$toString": "$" + key}, value}}}
			andCondition = append(andCondition, expr)
		}

		filter["$and"] = andCondition
	}

	return filter
}

func createQueryData(body []byte) (queryData, error) {
	var query queryData
	err := bson.UnmarshalExtJSON(body, true, &query)
	if err != nil {
		return queryData{}, err
	}

	if query.Find == nil {
		query.Find = bson.M{}
	}
	if query.Sort == nil {
		query.Sort = bson.M{}
	}
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PerPage == 0 {
		query.PerPage = 10
	}

	return query, nil
}

func createObjectID(_id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return [12]byte{}, err
	}

	return objectId, nil
}

func createDocument(body []byte) (interface{}, error) {
	var doc interface{}
	err := bson.UnmarshalExtJSON(body, true, &doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getIDResult(result *mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func createPaginateSkip(paginate queryData) int64 {
	var skip int64

	skip = 0

	if paginate.Page > 1 {
		skip = (paginate.Page * paginate.PerPage) - paginate.PerPage
	}

	return skip
}

func createPaginateTotalPages(perPage int64, count int64) int64 {
	div := float64(count) / float64(perPage)
	return int64(math.Ceil(div))
}
