// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/protodev-site/runtime/middleware"
)

// GetUserByNameHandlerFunc turns a function with the right signature into a get user by name handler
type GetUserByNameHandlerFunc func(GetUserByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserByNameHandlerFunc) Handle(params GetUserByNameParams) middleware.Responder {
	return fn(params)
}

// GetUserByNameHandler interface for that can handle valid get user by name params
type GetUserByNameHandler interface {
	Handle(GetUserByNameParams) middleware.Responder
}

// NewGetUserByName creates a new http.Handler for the get user by name operation
func NewGetUserByName(ctx *middleware.Context, handler GetUserByNameHandler) *GetUserByName {
	return &GetUserByName{Context: ctx, Handler: handler}
}

/* GetUserByName swagger:route GET /users/{username} user getUserByName

Get user by user name

*/
type GetUserByName struct {
	Context *middleware.Context
	Handler GetUserByNameHandler
}

func (o *GetUserByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserByNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
