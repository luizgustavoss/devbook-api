package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}