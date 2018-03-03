package router_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hendra-Huang/go-standard-layout/router"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestNew(t *testing.T) {
	options := router.Options{
		Timeout: 1 * time.Second,
	}
	tracer := mocktracer.New()
	rtr := router.New(options, tracer)
	testingutil.Assert(t, rtr != nil, "New returns nil")
}

func TestGet(t *testing.T) {
	options := router.Options{
		Timeout: 1 * time.Second,
	}
	tracer := mocktracer.New()
	rtr := router.New(options, tracer)
	rtr.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "testing")
	})

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	testingutil.Ok(t, err)
	body, err := ioutil.ReadAll(res.Body)
	testingutil.Ok(t, err)
	defer res.Body.Close()

	testingutil.Equals(t, http.StatusOK, res.StatusCode)
	testingutil.Equals(t, "testing", string(body))
}

func TestPost(t *testing.T) {
	options := router.Options{
		Timeout: 1 * time.Second,
	}
	tracer := mocktracer.New()
	rtr := router.New(options, tracer)
	rtr.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "testing")
	})

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", nil)
	testingutil.Ok(t, err)
	body, err := ioutil.ReadAll(res.Body)
	testingutil.Ok(t, err)
	defer res.Body.Close()

	testingutil.Equals(t, http.StatusOK, res.StatusCode)
	testingutil.Equals(t, "testing", string(body))
}

func TestGetWithMethodNotAllowed(t *testing.T) {
	options := router.Options{
		Timeout: 1 * time.Second,
	}
	tracer := mocktracer.New()
	rtr := router.New(options, tracer)
	rtr.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "testing")
	})

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	testingutil.Ok(t, err)

	testingutil.Equals(t, http.StatusMethodNotAllowed, res.StatusCode)
}

func TestGetWithTimeout(t *testing.T) {
	options := router.Options{
		Timeout: 10 * time.Millisecond,
	}
	tracer := mocktracer.New()
	rtr := router.New(options, tracer)
	rtr.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		fmt.Fprintf(w, "testing")
	})

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	testingutil.Ok(t, err)

	testingutil.Equals(t, http.StatusRequestTimeout, res.StatusCode)
}
