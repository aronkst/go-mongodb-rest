package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)

	body, err := getBody(r)
	if err != nil {
		httpError(w, err)
		return
	}

	data, err := s.MongoDB.Insert(collectionName, body)
	if err != nil {
		httpError(w, err)
		return
	}

	output, err := createOutput(data)
	if err != nil {
		httpError(w, err)
		return
	}

	httpSuccess(w, output, http.StatusCreated)
}
