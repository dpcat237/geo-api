package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"

	"gitlab.com/dpcat237/geoapi/model"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// getVariable returns request variable
func getVariable(r *http.Request, key string) string {
	return mux.Vars(r)[key]
}

// returnFailed replies to the request with the specified error message and HTTP code
func returnFailed(w http.ResponseWriter, er model.Error) {
	msg, err := er.String()
	if err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
	http.Error(w, msg, int(er.Status))
}

// returnFailed replies to the request with HTTP code 400
func returnBadRequest(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}

// returnFailed encode struct to JSON and replies it to the request
func returnJson(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}
