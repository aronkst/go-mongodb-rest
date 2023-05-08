package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Show(writer http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	id := getID(ps)

	data, err := s.MongoDB.Show(collectionName, id)
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
