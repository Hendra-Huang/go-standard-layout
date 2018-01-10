package router

import (
	"net/http"

	"github.com/Hendra-Huang/go-standard-layout/server/handler"
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterRoute(rtr *Router, pingHandler *handler.PingHandler, userHandler *handler.UserHandler) {
	rtr.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		prometheus.Handler()
	})
	rtr.Get("/ping", pingHandler.Ping)

	apiRouter := rtr.SubRouter("/api/v1")
	apiRouter.Get("/users", userHandler.GetAllUsers)
	apiRouter.Get("/user/{id}", userHandler.GetUserByID)
}
