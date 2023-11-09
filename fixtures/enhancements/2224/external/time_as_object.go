// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	timeext "time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/circl-dev/runtime"
)

// TimeAsObject time as object
//
// swagger:model TimeAsObject
type TimeAsObject struct {
	timeext.Time
}

func (m TimeAsObject) Validate(formats strfmt.Registry) error {
	var f interface{} = m.Time
	if v, ok := f.(runtime.Validatable); ok {
		return v.Validate(formats)
	}
	return nil
}

func (m TimeAsObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var f interface{} = m.Time
	if v, ok := f.(runtime.ContextValidatable); ok {
		return v.ContextValidate(ctx, formats)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TimeAsObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeAsObject) UnmarshalBinary(b []byte) error {
	var res TimeAsObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}