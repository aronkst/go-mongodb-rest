package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	_id := getID(ps)

	data, err := s.MongoDB.Show(collectionName, _id)
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
