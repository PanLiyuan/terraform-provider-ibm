// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumesClone volumes clone
// swagger:model VolumesClone
type VolumesClone struct {

	// Current action performed for the volumes-clone request
	Action string `json:"action,omitempty"`

	// The percent completion for the current action
	ActionPercentComplete int64 `json:"actionPercentComplete,omitempty"`

	// Creation Date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// Failure reason for a failed volumes-clone request
	FailureMessage string `json:"failureMessage,omitempty"`

	// Last Update Date
	// Format: date-time
	LastUpdateDate strfmt.DateTime `json:"lastUpdateDate,omitempty"`

	// Name assigned to a volumes-clone request
	Name string `json:"name,omitempty"`

	// Current status of the volumes-clone request
	Status string `json:"status,omitempty"`

	// ID assigned to a volumes-clone request
	VolumesCloneID string `json:"volumesCloneID,omitempty"`
}

// Validate validates this volumes clone
func (m *VolumesClone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumesClone) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VolumesClone) validateLastUpdateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumesClone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumesClone) UnmarshalBinary(b []byte) error {
	var res VolumesClone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
