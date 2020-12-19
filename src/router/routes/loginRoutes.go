package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var loginRoutes = []Route {
	{
		URI: "/login",
		Method: http.MethodPost,
		Function: controllers.Login,
		RequiresAuthentication: false,
	},
}


