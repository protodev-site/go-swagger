// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/protodev-site/runtime/middleware"
)

// GetAuthCallbackHandlerFunc turns a function with the right signature into a get auth callback handler
type GetAuthCallbackHandlerFunc func(GetAuthCallbackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAuthCallbackHandlerFunc) Handle(params GetAuthCallbackParams) middleware.Responder {
	return fn(params)
}

// GetAuthCallbackHandler interface for that can handle valid get auth callback params
type GetAuthCallbackHandler interface {
	Handle(GetAuthCallbackParams) middleware.Responder
}

// NewGetAuthCallback creates a new http.Handler for the get auth callback operation
func NewGetAuthCallback(ctx *middleware.Context, handler GetAuthCallbackHandler) *GetAuthCallback {
	return &GetAuthCallback{Context: ctx, Handler: handler}
}

/* GetAuthCallback swagger:route GET /auth/callback getAuthCallback

return access_token

*/
type GetAuthCallback struct {
	Context *middleware.Context
	Handler GetAuthCallbackHandler
}

func (o *GetAuthCallback) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAuthCallbackParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetAuthCallbackOKBody get auth callback o k body
//
// swagger:model GetAuthCallbackOKBody
type GetAuthCallbackOKBody struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`
}

// Validate validates this get auth callback o k body
func (o *GetAuthCallbackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get auth callback o k body based on context it is used
func (o *GetAuthCallbackOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAuthCallbackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAuthCallbackOKBody) UnmarshalBinary(b []byte) error {
	var res GetAuthCallbackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
