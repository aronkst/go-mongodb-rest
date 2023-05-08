package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) List(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	params := getQueryParams(request)

	data, err := s.MongoDB.List(collectionName, params)
	if err != nil {
		httpError(writer, err)

		return
	}

	output, err := createOutput(data)
	if err != nil {
		httpError(writer, err)

		return
	}

	httpSuccess(writer, output, http.StatusOK)
}
