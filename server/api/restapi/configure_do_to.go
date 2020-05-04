// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/rusakov/doto/server/service/goal_impl"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/rusakov/doto/server/api/restapi/operations"
	"github.com/rusakov/doto/server/api/restapi/operations/goal"
	"github.com/rusakov/doto/server/api/restapi/operations/task"
)

//go:generate swagger generate server --target ../go/src/github.com/rusakov/doto/server/api --name  --spec ../swagger.yaml

func configureFlags(api *operations.DoToAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DoToAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GoalGoalPostHandler = goal.GoalPostHandlerFunc(func(params goal.GoalPostParams) middleware.Responder {
		return middleware.NotImplemented("operation goal.GoalPost has not yet been implemented")
	})
	api.GoalGoalsSectionHandler = goal.GoalsSectionHandlerFunc(goal_impl.GoalsSection)
	/*func(params goal.GoalsSectionParams) middleware.Responder {
	//swagger.GoalsSection()
	return middleware.NotImplemented("operation goal.GoalsSection has not yet been implemented")
})*/
	api.GoalGoalsSectionQueryHandler = goal.GoalsSectionQueryHandlerFunc(func(params goal.GoalsSectionQueryParams) middleware.Responder {
		return middleware.NotImplemented("operation goal.GoalsSectionQuery has not yet been implemented")
	})
	api.TaskTaskPostHandler = task.TaskPostHandlerFunc(func(params task.TaskPostParams) middleware.Responder {
		return middleware.NotImplemented("operation task.TaskPost has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
