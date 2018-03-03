package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gorilla/mux"
)

// Router wrapper
type Router struct {
	options Options
	r       *mux.Router
	tracer  opentracing.Tracer
}

// Options for router
type Options struct {
	Timeout time.Duration
}

type statusCodeTracker struct {
	http.ResponseWriter
	status int
}

func (w *statusCodeTracker) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// New router
func New(opts Options, tracer opentracing.Tracer) *Router {
	muxRouter := mux.NewRouter()
	rtr := &Router{
		r:       muxRouter,
		options: opts,
		tracer:  tracer,
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

// trace middleware
func (rtr Router) trace(pattern string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := rtr.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := rtr.tracer.StartSpan(fmt.Sprintf("HTTP %s %s", r.Method, pattern), ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.String())
		ext.Component.Set(sp, "net/http")
		w = &statusCodeTracker{w, 200}
		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))

		h(w, r)

		ext.HTTPStatusCode.Set(sp, uint16(w.(*statusCodeTracker).status))
		sp.Finish()
	}
}

// SubRouter return a new Router with path prefix
func (rtr *Router) SubRouter(pathPrefix string) *Router {
	muxSubrouter := rtr.r.PathPrefix(pathPrefix).Subrouter()

	return &Router{
		r:       muxSubrouter,
		options: rtr.options,
		tracer:  rtr.tracer,
	}
}

// Get function
func (rtr *Router) Get(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.trace(pattern, rtr.timeout(h)))).Methods("GET")
}

// Post function
func (rtr *Router) Post(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.trace(pattern, rtr.timeout(h)))).Methods("POST")
}

// Put function
func (rtr *Router) Put(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.trace(pattern, rtr.timeout(h)))).Methods("PUT")
}

// Delete function
func (rtr *Router) Delete(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.trace(pattern, rtr.timeout(h)))).Methods("DELETE")
}

// Patch function
func (rtr *Router) Patch(pattern string, h http.HandlerFunc) {
	rtr.r.HandleFunc(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.trace(pattern, rtr.timeout(h)))).Methods("PATCH")
}

// ServeHTTP function
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.r.ServeHTTP(w, r)
}
