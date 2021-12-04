package api

import (
	"net/http"
	"strings"
)

type DataStore interface {
	AddItemToStore(key string, value string)
	GetItemFromStore(key string) (string, error)
}

type ApiServer struct {
	store DataStore
}

func NewApiServer(s DataStore) ApiServer {
	return ApiServer{store: s}
}

func (a *ApiServer) GetItem() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/getvalue/")
		dt, err := a.store.GetItemFromStore(key)

		if err != nil {
			ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		ResponseWithJSON(w, http.StatusOK, dt)
	}
}

func (a *ApiServer) AddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		key := r.Form.Get("key")
		value := r.Form.Get("value")

		a.store.AddItemToStore(key, value)

		ResponseWithJSON(w, http.StatusAccepted, nil)
	}
}
