/*
 * Do-to
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	cors.Default().Handler(router)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	/*Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"GoalPost",
		strings.ToUpper("Post"),
		"/api/goal/{section}",
		GoalPost,
	},

	Route{
		"GoalsSection",
		strings.ToUpper("Get"),
		"/api/goals/{section}",
		GoalsSection,
	},

	Route{
		"GoalsSectionQuery",
		strings.ToUpper("Get"),
		"/api/goals",
		GoalsQuery,
	},*/
}
