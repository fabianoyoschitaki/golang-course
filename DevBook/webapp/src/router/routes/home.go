package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homepageRoute = []Route{
	{
		URI:                    "/home",
		Method:                 http.MethodGet,
		Function:               controllers.LoadHomepage,
		RequiresAuthentication: true,
	},
}
