package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"erros"`
}

func ResponseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

func ResponseError(w http.ResponseWriter, statusCode int, errMessage string) {
	// ResponseWithJSON(w, statusCode, map[string]string{"Error": errMessage})
	ResponseWithJSON(w, statusCode, ErrorResponse{Code: statusCode, Status: "Error", Message: errMessage})
}

// func logMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }