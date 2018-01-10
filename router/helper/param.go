package helper

import (
	"net/http"

	"github.com/gorilla/mux"
)

// URLParam get param from rest request
func URLParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	return params[key]
}
