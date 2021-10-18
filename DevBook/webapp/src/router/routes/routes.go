package routes

import (
	"log"
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents al web routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Configure adds all routes inside router
func Configure(router *mux.Router) *mux.Router {

	routes := loginRoutes                     // login routes
	routes = append(routes, userRoutes...)    // users routes
	routes = append(routes, homepageRoute...) // home page route
	routes = append(routes, postsRoutes...)   // posts route

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
		log.Printf("Route %s %s auth: %t configured\n", route.Method, route.URI, route.RequiresAuthentication)
	}

	// we need to serve our css and js files
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer)) // instead of using ../assets inside our html files, we just provide /assets

	return router
}
