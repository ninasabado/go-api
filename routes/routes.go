package routes

import (
	"../handles"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*
   Define standardized structs, including a Route
*/
type (
	Route struct {
		Path   string
		Method string
		Handle httprouter.Handle
	}
)

/*
   Source of all the Routes
*/
func SetRoutes() *httprouter.Router {

	// Instantiates a router with all defaults enabled
	router := httprouter.New()

	// TODO: Pass in defined NotFound, MethodNotAllowed and PanicHandler Handlers
	// router.NotFound = some handler
	// router.MethodNotAllowed = some handler
	// router.PanicHandler = func(http.ResponseWriter, *http.Request, interface{})

	setRoutes([]Route{
		Route{"/", http.MethodGet, handles.Index},
		Route{"/hello/:name", http.MethodGet, handles.Hello},
	}, router)

	return router
}

/*
   Helper functions
*/
func setRoutes(routes []Route, router *httprouter.Router) {

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Path, route.Handle)
		case http.MethodPut:
			router.PUT(route.Path, route.Handle)
		case http.MethodPatch:
			router.PATCH(route.Path, route.Handle)
		case http.MethodDelete:
			router.DELETE(route.Path, route.Handle)
		}
	}
}
