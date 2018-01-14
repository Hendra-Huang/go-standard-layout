package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Hendra-Huang/go-standard-layout/server"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestNew(t *testing.T) {
	srv := server.New(server.Options{
		ListenAddress: ":9090",
	})
	testingutil.Assert(t, srv != nil, "Server is nil")
}

func TestServe(t *testing.T) {
	srv := server.New(server.Options{
		ListenAddress: ":9090",
	})
	mux := http.DefaultServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "testing")
	})
	go func() {
		err := srv.Serve(mux)
		testingutil.Ok(t, err)
	}()
	time.Sleep(1 * time.Second)

	res, err := http.Get("http://localhost:9090/")
	testingutil.Ok(t, err)
	body, err := ioutil.ReadAll(res.Body)
	testingutil.Ok(t, err)
	defer res.Body.Close()

	testingutil.Equals(t, http.StatusOK, res.StatusCode)
	testingutil.Equals(t, "testing", string(body))
}
