package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	params := getQueryParams(r)

	data, err := s.MongoDB.List(collectionName, params)
	if err != nil {
		httpError(w, err)
		return
	}

	output, err := createOutput(data)
	if err != nil {
		httpError(w, err)
		return
	}

	httpSuccess(w, output, http.StatusOK)
}
