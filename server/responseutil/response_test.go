package responseutil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/server/responseutil"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestWithHeader(t *testing.T) {
	writerOption := responseutil.WithHeader("Content-Type", "plain/text")
	rr := httptest.NewRecorder()
	writerOption(rr)

	testingutil.Equals(t, "plain/text", rr.Header().Get("Content-Type"))
}

func TestWithStatus(t *testing.T) {
	writerOption := responseutil.WithStatus(http.StatusPermanentRedirect)
	rr := httptest.NewRecorder()
	writerOption(rr)

	testingutil.Equals(t, http.StatusPermanentRedirect, rr.Code)
}

func TestBadRequest(t *testing.T) {
	message := "id is invalid"
	rr := httptest.NewRecorder()
	responseutil.BadRequest(rr, message)
	expectedMessage := message + "\n"

	testingutil.Equals(t, http.StatusBadRequest, rr.Code)
	testingutil.Equals(t, expectedMessage, rr.Body.String())
}

func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	responseutil.NotFound(rr)

	testingutil.Equals(t, http.StatusNotFound, rr.Code)
}

func TestInternalServerError(t *testing.T) {
	rr := httptest.NewRecorder()
	responseutil.InternalServerError(rr)

	testingutil.Equals(t, http.StatusInternalServerError, rr.Code)
}
