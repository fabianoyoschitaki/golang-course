package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	// gets our mux router configured
	return routes.Configure(r)
}
