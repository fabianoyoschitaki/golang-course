package routes

import (
	"net/http"
	"webapp/src/controllers"
)

// either / or /login points to login
var loginRoutes = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodPost,
		Function:               controllers.AttemptLogin,
		RequiresAuthentication: false,
	},
}
