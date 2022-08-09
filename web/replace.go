package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Replace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	params := getQueryParams(r)

	body, err := getBody(r)
	if err != nil {
		httpError(w, err)
		return
	}

	data, err := s.MongoDB.Replace(collectionName, params, body)
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
