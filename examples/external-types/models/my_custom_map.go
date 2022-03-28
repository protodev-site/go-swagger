// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	alternate "github.com/circl-dev/go-swagger/examples/external-types/fred"
)

// MyCustomMap This generate a map type in models, based on the external type.
//
// MyCustomMap map[string]map[string]alternate.MyAlternateType
//
// The validation method of the external type is called by the generated map.
//
//
// swagger:model MyCustomMap
type MyCustomMap map[string]map[string]alternate.MyAlternateType

// Validate validates this my custom map
func (m MyCustomMap) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			if val, ok := m[k][kk]; ok {
				if err := val.Validate(formats); err != nil {
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this my custom map based on context it is used
func (m MyCustomMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
