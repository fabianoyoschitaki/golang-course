package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a Router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()

	// gets our mux router configured
	return routes.Configure(r)
}
