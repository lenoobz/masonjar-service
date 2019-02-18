/**
 * Create by Le Trong on 17/Feb/2019
 */

package service

import "github.com/gorilla/mux"

// NewRouter - Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {
		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
