package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Insert(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)

	body, err := getBody(request)
	if err != nil {
		httpError(writer, err)

		return
	}

	data, err := s.MongoDB.Insert(collectionName, body)
	if err != nil {
		httpError(writer, err)

		return
	}

	output, err := createOutput(data)
	if err != nil {
		httpError(writer, err)

		return
	}

	httpSuccess(writer, output, http.StatusCreated)
}
