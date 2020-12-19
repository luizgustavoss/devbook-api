package routes

import (
	"devbook/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// represents API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// configures routes in router
func ConfigureRoutes(r *mux.Router) *mux.Router {

	routes := userRoutes
	routes =  append(routes, loginRoutes...)

	for _, route := range routes {

		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middlewares.CheckAuthenticatedRequest(
					middlewares.LogRequest(route.Function))).Methods(route.Method)
		} else{
			r.HandleFunc(route.URI,
				middlewares.LogRequest(route.Function)).Methods(route.Method)
		}
	}
	return r
}