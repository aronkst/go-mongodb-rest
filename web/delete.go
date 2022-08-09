package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	params := getQueryParams(r)

	err := s.MongoDB.Delete(collectionName, params)
	if err != nil {
		httpError(w, err)
		return
	}

	httpNoContent(w)
}
