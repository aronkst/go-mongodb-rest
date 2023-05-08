package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Paginate(writer http.ResponseWriter, equest *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)

	body, err := getBody(equest)
	if err != nil {
		httpError(writer, err)

		return
	}

	data, err := s.MongoDB.Paginate(collectionName, body)
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
