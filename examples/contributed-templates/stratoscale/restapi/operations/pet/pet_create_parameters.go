// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/circl-dev/runtime"
	"github.com/circl-dev/runtime/middleware"
	"github.com/circl-dev/validate"

	"github.com/circl-dev/go-swagger/examples/contributed-templates/stratoscale/models"
)

// NewPetCreateParams creates a new PetCreateParams object
//
// There are no default values defined in the spec.
func NewPetCreateParams() PetCreateParams {

	return PetCreateParams{}
}

// PetCreateParams contains all the bound params for the pet create operation
// typically these are obtained from a http.Request
//
// swagger:parameters PetCreate
type PetCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Pet object that needs to be added to the store
	  Required: true
	  In: body
	*/
	Body *models.Pet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPetCreateParams() beforehand.
func (o *PetCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Pet
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
