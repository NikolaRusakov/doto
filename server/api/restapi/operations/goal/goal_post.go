// Code generated by go-swagger; DO NOT EDIT.

package goal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GoalPostHandlerFunc turns a function with the right signature into a goal post handler
type GoalPostHandlerFunc func(GoalPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GoalPostHandlerFunc) Handle(params GoalPostParams) middleware.Responder {
	return fn(params)
}

// GoalPostHandler interface for that can handle valid goal post params
type GoalPostHandler interface {
	Handle(GoalPostParams) middleware.Responder
}

// NewGoalPost creates a new http.Handler for the goal post operation
func NewGoalPost(ctx *middleware.Context, handler GoalPostHandler) *GoalPost {
	return &GoalPost{Context: ctx, Handler: handler}
}

/*GoalPost swagger:route POST /goal/{section} Goal goalPost

post goal

Adds goal

*/
type GoalPost struct {
	Context *middleware.Context
	Handler GoalPostHandler
}

func (o *GoalPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGoalPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
