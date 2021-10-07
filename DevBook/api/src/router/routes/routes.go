package routes

import (
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

	// routes from users.go
	routes := routeUsers
	// route from login.go
	routes = append(routes, loginRoute)

	// a simple for we can configure all routes for our application
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
