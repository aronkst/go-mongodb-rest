package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aronkst/go-mongodb-rest/mongodb"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	MongoDB *mongodb.MongoDB
}

func New(mongoDB *mongodb.MongoDB) *Server {
	return &Server{MongoDB: mongoDB}
}

func httpSuccess(writer http.ResponseWriter, output []byte, httpStatus int) {
	writer.WriteHeader(httpStatus)
	writer.Header().Set("Content-Type", "application/json")

	if _, err := writer.Write(output); err != nil {
		log.Fatal(err)
	}
}

func httpNoContent(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNoContent)
}

func httpError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}

func getCollectionName(ps httprouter.Params) string {
	return ps.ByName("collection")
}

func getID(ps httprouter.Params) string {
	return ps.ByName("_id")
}

func getBody(request *http.Request) ([]byte, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getQueryParams(request *http.Request) map[string]string {
	params := make(map[string]string)

	query := request.URL.Query()
	for key := range query {
		params[key] = query.Get(key)
	}

	return params
}

func createOutput(data any) ([]byte, error) {
	output, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
