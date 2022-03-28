// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime/middleware"
)

// NewDeleteOrderParams creates a new DeleteOrderParams object
//
// There are no default values defined in the spec.
func NewDeleteOrderParams() DeleteOrderParams {

	return DeleteOrderParams{}
}

// DeleteOrderParams contains all the bound params for the delete order operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteOrder
type DeleteOrderParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of the order that needs to be deleted
	  Required: true
	  In: path
	*/
	OrderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteOrderParams() beforehand.
func (o *DeleteOrderParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrderID, rhkOrderID, _ := route.Params.GetOK("orderId")
	if err := o.bindOrderID(rOrderID, rhkOrderID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrderID binds and validates parameter OrderID from path.
func (o *DeleteOrderParams) bindOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.OrderID = raw

	return nil
}
