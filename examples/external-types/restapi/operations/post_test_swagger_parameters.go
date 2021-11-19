// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/protodev-site/runtime"
	"github.com/protodev-site/runtime/middleware"

	custom "github.com/protodev-site/go-swagger/examples/external-types/fred"
)

// NewPostTestParams creates a new PostTestParams object
//
// There are no default values defined in the spec.
func NewPostTestParams() PostTestParams {

	return PostTestParams{}
}

// PostTestParams contains all the bound params for the post test operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostTest
type PostTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Defines a parameter as an array of external types.
	The body parameter is defined as []custom.MyAlternateString

	No definition is generated in models.

	  In: body
	*/
	CustomizedStrings []custom.MyAlternateString
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostTestParams() beforehand.
func (o *PostTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []custom.MyAlternateString
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("customizedStrings", "body", "", err))
		} else {

			// validate array of body objects
			for i := range body {
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.CustomizedStrings = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
