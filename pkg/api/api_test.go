package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubDataStore struct {
	store map[string]string
}

func (s *StubDataStore) AddItemToStore(key string, value string) {
	s.store[key] = value
}

func (s *StubDataStore) GetItemFromStore(key string) (string, error) {
	return s.store[key], nil
}

func (s *StubDataStore) FlushStore() {
	s.store = make(map[string]string)
}

func Test_GetValue_StatusOK_Response(t *testing.T) {
	store := StubDataStore{
		map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/getvalue?key=key1", nil)
		response := httptest.NewRecorder()

		server.GetValue().ServeHTTP(response, request)

		assertStatus(t, http.StatusOK, response.Code)
		assertResponseBody(t, "value1", store.store["key1"])

	})
}

func Test_GetValue_StatusNotFound_KeyParameter(t *testing.T) {
	store := StubDataStore{
		map[string]string{},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {
		expectedResponseBody := ErrorResponse{
			Code:    404,
			Status:  "Error",
			Message: "Key parameter not Found in URL",
		}
		request := httptest.NewRequest(http.MethodGet, "/getvalue?key=", nil)
		response := httptest.NewRecorder()

		server.GetValue().ServeHTTP(response, request)

		assertStatus(t, http.StatusNotFound, response.Code)

		//want, _ := json.Marshal(expectedResponseBody)
		var got ErrorResponse
		json.NewDecoder(response.Body).Decode(&got)
		assertResponseBody(t, expectedResponseBody, got)

	})
}
func Test_AddValue_StatusOK(t *testing.T) {
	store := StubDataStore{
		map[string]string{},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {
		postBody := &PostBody{
			Key:   "test1",
			Value: "test11",
		}
		postBodyJson, _ := json.Marshal(postBody)

		request := httptest.NewRequest(http.MethodPost, "/setvalue", strings.NewReader(string(postBodyJson)))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.AddValue().ServeHTTP(response, request)

		assertStatus(t, http.StatusOK, response.Code)

		if len(store.store) != 1 {
			t.Fatalf("After setValue call, store count not correct got %d and want %d", len(store.store), 1)
		}

	})
}

func Test_AddValue_StatusBadRequest(t *testing.T) {
	store := StubDataStore{
		map[string]string{},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {

		postBodyJson := "{\"Key\":\"test1\"\"Value\":\"test11\"}"

		request := httptest.NewRequest(http.MethodPost, "/setvalue", strings.NewReader(string(postBodyJson)))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.AddValue().ServeHTTP(response, request)

		assertStatus(t, http.StatusBadRequest, response.Code)

	})
}

func Test_Flush_StatusCode_NoContent(t *testing.T) {
	store := StubDataStore{
		map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/flush", nil)
		response := httptest.NewRecorder()

		server.Flush().ServeHTTP(response, request)

		assertStatus(t, http.StatusNoContent, response.Code)

		if len(store.store) != 0 {
			t.Errorf("Not correct flush, store count %d, should be %d", len(store.store), 0)
		}
	})
}

func Test_Flush_StoreLen_Zero(t *testing.T) {
	store := StubDataStore{
		map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	server := &ApiServer{&store}

	t.Run(t.Name(), func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/flush", nil)
		response := httptest.NewRecorder()

		server.Flush().ServeHTTP(response, request)

		if len(store.store) != 0 {
			t.Errorf("Not correct flush, store count %d, should be %d", len(store.store), 0)
		}
	})
}

func assertStatus(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("Not correct status code, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, want, got interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Not correct response body, got %q, want %q", got, want)
	}
}
