// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/protodev-site/validate"
)

// Milestone A milestone is a particular goal that is important to the project for this issue tracker.
//
// Milestones can have a escription and due date.
// This can be useful for filters and such.
//
//
// swagger:model Milestone
type Milestone struct {

	// The description of the milestone.
	//
	// A description is a free text field that allows for a more detailed explanation of what the milestone is trying to achieve.
	//
	Description string `json:"description,omitempty"`

	// An optional due date for this milestone.
	//
	// This property is optional, but when present it lets people know when they can expect this milestone to be completed.
	//
	// Format: date
	DueDate strfmt.Date `json:"dueDate,omitempty"`

	// The name of the milestone.
	//
	// Each milestone should get a unique name.
	//
	// Required: true
	// Max Length: 50
	// Min Length: 3
	// Pattern: [A-Za-z][\w- ]+
	Name *string `json:"name"`

	// stats
	Stats *MilestoneStats `json:"stats,omitempty"`
}

// Validate validates this milestone
func (m *Milestone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Milestone) validateDueDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Milestone) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[A-Za-z][\w- ]+`); err != nil {
		return err
	}

	return nil
}

func (m *Milestone) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this milestone based on the context it is used
func (m *Milestone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Milestone) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {
		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Milestone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Milestone) UnmarshalBinary(b []byte) error {
	var res Milestone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MilestoneStats Some counters for this milestone.
//
// This object contains counts for the remaining open issues and the amount of issues that have been closed.
//
//
// swagger:model MilestoneStats
type MilestoneStats struct {

	// The closed issues.
	Closed int32 `json:"closed,omitempty"`

	// The remaining open issues.
	Open int32 `json:"open,omitempty"`

	// The total number of issues for this milestone.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this milestone stats
func (m *MilestoneStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this milestone stats based on context it is used
func (m *MilestoneStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MilestoneStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MilestoneStats) UnmarshalBinary(b []byte) error {
	var res MilestoneStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
