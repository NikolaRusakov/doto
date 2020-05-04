// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Task task
// swagger:model Task
type Task struct {

	// description
	Description string `json:"description,omitempty"`

	// estimation
	Estimation string `json:"estimation,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// relationships
	Relationships []*Relationship `json:"relationships"`

	// section
	Section GoalSections `json:"section,omitempty"`

	// subtask
	Subtask *Task `json:"subtask,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	for i := 0; i < len(m.Relationships); i++ {
		if swag.IsZero(m.Relationships[i]) { // not required
			continue
		}

		if m.Relationships[i] != nil {
			if err := m.Relationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateSection(formats strfmt.Registry) error {

	if swag.IsZero(m.Section) { // not required
		return nil
	}

	if err := m.Section.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("section")
		}
		return err
	}

	return nil
}

func (m *Task) validateSubtask(formats strfmt.Registry) error {

	if swag.IsZero(m.Subtask) { // not required
		return nil
	}

	if m.Subtask != nil {
		if err := m.Subtask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subtask")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}