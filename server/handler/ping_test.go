package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/server/handler"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestNewPingHandler(t *testing.T) {
	pingHandler := handler.NewPingHandler()
	testingutil.Assert(t, pingHandler != nil, "PingHandler is nil")
}

func TestPing(t *testing.T) {
	pingHandler := handler.NewPingHandler()

	req, err := http.NewRequest("GET", "/ping", nil)
	testingutil.Ok(t, err)
	rr := httptest.NewRecorder()
	httphandler := http.HandlerFunc(pingHandler.Ping)
	httphandler.ServeHTTP(rr, req)

	testingutil.Equals(t, http.StatusOK, rr.Code)
	testingutil.Equals(t, "pong", rr.Body.String())
}
