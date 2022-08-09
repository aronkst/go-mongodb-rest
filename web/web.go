package web

import (
	"encoding/json"
	"io/ioutil"
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

func httpSuccess(w http.ResponseWriter, output []byte, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func httpNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func httpError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func getCollectionName(ps httprouter.Params) string {
	return ps.ByName("collection")
}

func getID(ps httprouter.Params) string {
	return ps.ByName("_id")
}

func getBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getQueryParams(r *http.Request) map[string]string {
	params := make(map[string]string)

	query := r.URL.Query()
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
