package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route has URI, Method, the function to deal with it and if it requires authentication
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Configure configures all routes
func Configure(r *mux.Router) *mux.Router {

	routes := routeUsers                // routes from users.go
	routes = append(routes, loginRoute) // route from login.go

	// a simple for we can configure all routes for our application
	for _, route := range routes {

		loggingRoute := middlewares.Logger(route.Function)
		// if it requires authentication, we run the authenticate function before the actual method
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Authenticate(loggingRoute)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, loggingRoute).Methods(route.Method)
		}
	}
	return r
}
