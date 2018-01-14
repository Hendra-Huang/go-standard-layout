package responseutil

import (
	"encoding/json"
	"net/http"

	"github.com/Hendra-Huang/go-standard-layout/errorutil"
	"github.com/Hendra-Huang/go-standard-layout/log"
)

// JSON sets the response as json
func JSON(w http.ResponseWriter, data interface{}, options ...WriterOption) {
	w.Header().Set("Content-Type", "application/json")
	for _, option := range options {
		option(w)
	}
	body, err := json.Marshal(data)
	errorutil.CheckWithErrorHandler(err, func(err error) {
		log.Error(err)
		InternalServerError(w)
	})
	w.Write(body)
}
