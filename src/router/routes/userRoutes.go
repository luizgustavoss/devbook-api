package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var userRoutes = []Route {
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.ListUsers,
		RequiresAuthentication: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: controllers.FindUserById,
		RequiresAuthentication: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
