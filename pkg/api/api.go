package api

import (
	"encoding/json"
	"net/http"
)

type DataStore interface {
	AddItemToStore(key string, value string)
	GetItemFromStore(key string) (string, error)
}

type ApiServer struct {
	store DataStore
}

type PostBody struct {
	Key   string
	Value string
}

func NewApiServer(s DataStore) ApiServer {
	return ApiServer{store: s}
}

func (a *ApiServer) GetItem() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//key := strings.TrimPrefix(r.URL.Path, "/getvalue/")
		keys := r.URL.Query()["key"]

		if len(keys[0]) < 1 {
			ResponseError(w, http.StatusNotFound, "error")
			return
		}

		key := keys[0]

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

		decoder := json.NewDecoder(r.Body)

		var body PostBody

		err := decoder.Decode(&body)

		if err != nil {
			ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		// r.ParseForm()
		// key := r.Form.Get("key")
		// value := r.Form.Get("value")

		a.store.AddItemToStore(body.Key, body.Value)

		ResponseWithJSON(w, http.StatusOK, nil) //statusaccepted??
	}
}
