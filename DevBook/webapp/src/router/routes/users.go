package routes

import (
	"net/http"
	"webapp/src/controllers"
)

// either / or /login points to login
var userRoutes = []Route{
	{
		URI:                    "/signup",
		Method:                 http.MethodGet,
		Function:               controllers.LoadSignUpPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
}
