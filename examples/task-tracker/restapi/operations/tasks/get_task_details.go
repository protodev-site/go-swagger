// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/protodev-site/runtime/middleware"
)

// GetTaskDetailsHandlerFunc turns a function with the right signature into a get task details handler
type GetTaskDetailsHandlerFunc func(GetTaskDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskDetailsHandlerFunc) Handle(params GetTaskDetailsParams) middleware.Responder {
	return fn(params)
}

// GetTaskDetailsHandler interface for that can handle valid get task details params
type GetTaskDetailsHandler interface {
	Handle(GetTaskDetailsParams) middleware.Responder
}

// NewGetTaskDetails creates a new http.Handler for the get task details operation
func NewGetTaskDetails(ctx *middleware.Context, handler GetTaskDetailsHandler) *GetTaskDetails {
	return &GetTaskDetails{Context: ctx, Handler: handler}
}

/* GetTaskDetails swagger:route GET /tasks/{id} tasks getTaskDetails

Gets the details for a task.

The details view has more information than the card view.
You can see who reported the issue and who last updated it when.

There are also comments for each issue.


*/
type GetTaskDetails struct {
	Context *middleware.Context
	Handler GetTaskDetailsHandler
}

func (o *GetTaskDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTaskDetailsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
