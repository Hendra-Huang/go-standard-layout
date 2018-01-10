package responseutil

import "net/http"

// WriterOption for modifying the response
type WriterOption func(http.ResponseWriter)

// WithHeader add header to the response
func WithHeader(header string, value string) WriterOption {
	return func(w http.ResponseWriter) {
		w.Header().Set(header, value)
	}
}

// WithStatus define the response status code
// NOTE: need to be passed as the last argument because this function calls WriteHeader
func WithStatus(code int) WriterOption {
	return func(w http.ResponseWriter) {
		w.WriteHeader(code)
	}
}

// BadRequest response
func BadRequest(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

// NotFound response
func NotFound(w http.ResponseWriter) {
	http.Error(w, "", http.StatusNotFound)
}

// InternalServerError response
func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "", http.StatusInternalServerError)
}
