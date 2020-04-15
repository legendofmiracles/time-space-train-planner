// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StationDetails Additional details for a station that go beyond what is present in StationData or in a Timetable. The details include all station based messages.
//
// swagger:model stationDetails
type StationDetails struct {

	// EVA station number.
	// Required: true
	Eva *int64 `json:"eva" xml:"eva,attr"`

	// List of station based messages.
	M []*Message `json:"m" xml:"m"`
}

// Validate validates this station details
func (m *StationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEva(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StationDetails) validateEva(formats strfmt.Registry) error {

	if err := validate.Required("eva", "body", m.Eva); err != nil {
		return err
	}

	return nil
}

func (m *StationDetails) validateM(formats strfmt.Registry) error {

	if swag.IsZero(m.M) { // not required
		return nil
	}

	for i := 0; i < len(m.M); i++ {
		if swag.IsZero(m.M[i]) { // not required
			continue
		}

		if m.M[i] != nil {
			if err := m.M[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("m" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StationDetails) UnmarshalBinary(b []byte) error {
	var res StationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}