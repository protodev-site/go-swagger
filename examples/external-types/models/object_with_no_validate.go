// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	httpext "net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/circl-dev/validate"
)

// ObjectWithNoValidate A reference to the NoValidateExternal external type.
//
// If the "noValidation" hint is omitted in the definition above, this code won't build because `http.Request` has no `Validate` method.
//
//
// swagger:model ObjectWithNoValidate
type ObjectWithNoValidate struct {

	// my mandatory request
	// Required: true
	MyMandatoryRequest *httpext.Request `json:"myMandatoryRequest"`

	// my request
	MyRequest httpext.Request `json:"myRequest,omitempty"`
}

// Validate validates this object with no validate
func (m *ObjectWithNoValidate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMyMandatoryRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectWithNoValidate) validateMyMandatoryRequest(formats strfmt.Registry) error {

	if err := validate.Required("myMandatoryRequest", "body", m.MyMandatoryRequest); err != nil {
		return err
	}

	return nil
}

func (m *ObjectWithNoValidate) validateMyRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.MyRequest) { // not required
		return nil
	}

	return nil
}

// ContextValidate validates this object with no validate based on context it is used
func (m *ObjectWithNoValidate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectWithNoValidate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectWithNoValidate) UnmarshalBinary(b []byte) error {
	var res ObjectWithNoValidate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
