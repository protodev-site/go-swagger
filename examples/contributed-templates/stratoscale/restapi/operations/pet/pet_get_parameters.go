// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/circl-dev/runtime/middleware"
)

// NewPetGetParams creates a new PetGetParams object
//
// There are no default values defined in the spec.
func NewPetGetParams() PetGetParams {

	return PetGetParams{}
}

// PetGetParams contains all the bound params for the pet get operation
// typically these are obtained from a http.Request
//
// swagger:parameters PetGet
type PetGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of pet to return
	  Required: true
	  In: path
	*/
	PetID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPetGetParams() beforehand.
func (o *PetGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPetID, rhkPetID, _ := route.Params.GetOK("petId")
	if err := o.bindPetID(rPetID, rhkPetID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPetID binds and validates parameter PetID from path.
func (o *PetGetParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("petId", "path", "int64", raw)
	}
	o.PetID = value

	return nil
}
