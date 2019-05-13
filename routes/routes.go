package routes

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)


/*
    Define standardized structs, including a Route
*/
type (
    Route struct {
        Path        string
        Method      string
        Handle      httprouter.Handle
    }
)

/*
    Define methods that the API supports
*/
const (
    GET    = "GET"
    POST   = "POST"
    PATCH  = "PATCH"
    DELETE = "DELETE"
)

/*
    Source of all the Routes
*/
func SetRoutes() (*httprouter.Router){

    router := httprouter.New()

    setRoutes([]Route{
        Route{"/", GET, Index},
        Route{"/hello/:name", GET, Hello},
    }, router)

    // TODO: Check if any errors occur and bubble it up accordingly
    return router
}

/* 
    Sample function handlers
    TODO: Eventually have the source be different packages
 */
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

/*
    Helper functions
*/
func setRoutes(routes []Route, router *httprouter.Router){

    for _, route := range routes {
        if route.Method == GET {
            router.GET(route.Path, route.Handle)
        }
    }
}