package api

import (
	"encoding/json"
	"net/http"
)

type DataStore interface {
	AddItemToStore(key string, value string)
	GetItemFromStore(key string) string
	FlushStore()
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

func (a *ApiServer) GetValue() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//key := strings.TrimPrefix(r.URL.Path, "/getvalue/")
		keys := r.URL.Query()["key"]

		if len(keys[0]) < 1 {
			ResponseError(w, http.StatusNotFound, "Key parameter not Found in URL")
			return
		}

		key := keys[0]

		dt := a.store.GetItemFromStore(key)

		if dt == "" {
			ResponseError(w, http.StatusNotFound, key+" does not found in the storage")
			return
		}

		ResponseWithJSON(w, http.StatusOK, dt)
	}
}

func (a *ApiServer) AddValue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var body PostBody

		err := decoder.Decode(&body)

		if err != nil {
			ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		a.store.AddItemToStore(body.Key, body.Value)

		ResponseWithJSON(w, http.StatusOK, body) //statusaccepted??
	}
}

func (a *ApiServer) Flush() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		a.store.FlushStore()

		ResponseWithJSON(w, http.StatusNoContent, nil)
	}
}

func (a *ApiServer) Home() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		homeBody := "Depot-in-go RESTApi is running"
		ResponseWithJSON(w, http.StatusAccepted, homeBody)
	}
}
