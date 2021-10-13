package routes

import (
	"fmt"
	"net/http"

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

	routes := loginRoutes                  // login routes
	routes = append(routes, userRoutes...) // users routes

	for _, route := range routes {
		fmt.Printf("Route %s configured\n", route.URI)
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	// we need to serve our css and js files
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer)) // instead of using ../assets inside our html files, we just provide /assets

	return router
}
