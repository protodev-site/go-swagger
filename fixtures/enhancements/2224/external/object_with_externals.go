// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	httpext "net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectWithExternals object with externals
//
// swagger:model ObjectWithExternals
type ObjectWithExternals struct {

	// a
	A httpext.Request `json:"a,omitempty"`

	// b
	B httpext.Request `json:"b,omitempty"`
}

// Validate validates this object with externals
func (m *ObjectWithExternals) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateB(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectWithExternals) validateA(formats strfmt.Registry) error {
	if swag.IsZero(m.A) { // not required
		return nil
	}

	return nil
}

func (m *ObjectWithExternals) validateB(formats strfmt.Registry) error {
	if swag.IsZero(m.B) { // not required
		return nil
	}

	return nil
}

// ContextValidate validates this object with externals based on context it is used
func (m *ObjectWithExternals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectWithExternals) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectWithExternals) UnmarshalBinary(b []byte) error {
	var res ObjectWithExternals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}