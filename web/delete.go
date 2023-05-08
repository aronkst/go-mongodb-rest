package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Delete(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	collectionName := getCollectionName(ps)
	params := getQueryParams(request)

	err := s.MongoDB.Delete(collectionName, params)
	if err != nil {
		httpError(writer, err)

		return
	}

	httpNoContent(writer)
}
