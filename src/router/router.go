package router

import (
	"devbook/src/router/routes"
	"github.com/gorilla/mux"
)

// return a router with configured routes
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigureRoutes(router)
}