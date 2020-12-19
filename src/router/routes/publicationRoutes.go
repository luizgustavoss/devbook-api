package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/publications",
		Method:                 http.MethodGet,
		Function:               controllers.GetUserPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePublication,
		RequiresAuthentication: true,
	},

}
