package router

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gorilla/mux"
)

// Router wrapper
type Router struct {
	options Options
	r       *mux.Router
}

// Options for router
type Options struct {
	Timeout time.Duration
}

// New router
func New(opts Options) *Router {
	muxRouter := mux.NewRouter()
	rtr := &Router{
		r:       muxRouter,
		options: opts,
	}

	return rtr
}

// timeout middleware
func (rtr Router) timeout(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opts := rtr.options
		// cancel context
		if opts.Timeout > 0 {
			ctx, cancel := context.WithTimeout(r.Context(), opts.Timeout)
			defer cancel()
			r = r.WithContext(ctx)
		}

		doneChan := make(chan bool)
		go func() {
			h(w, r)
			doneChan <- true
		}()
		select {
		case <-r.Context().Done():
			w.WriteHeader(http.StatusRequestTimeout)
			return
		case <-doneChan:
			return
		}
	}
}

// SubRouter return a new Router with path prefix
func (rtr *Router) SubRouter(pathPrefix string) *Router {
	muxSubrouter := rtr.r.PathPrefix(pathPrefix).Subrouter()

	return &Router{
		r:       muxSubrouter,
		options: rtr.options,
	}
}

// Get function
func (rtr *Router) Get(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h))).Methods("GET")
}

// Post function
func (rtr *Router) Post(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h))).Methods("POST")
}

// Put function
func (rtr *Router) Put(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h))).Methods("PUT")
}

// Delete function
func (rtr *Router) Delete(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h))).Methods("DELETE")
}

// Patch function
func (rtr *Router) Patch(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h))).Methods("PATCH")
}

// ServeHTTP function
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.r.ServeHTTP(w, r)
}
