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

// TripStop This element extends Stop.
//
// swagger:model tripStop
type TripStop struct {

	// Arrival element. This element does not have child elements. All information about the arrival is stored in attributes (see the next table).
	Ar *Event `json:"ar,omitempty" xml:"ar,omitempty"`

	// Connection element.
	Conn []*Connection `json:"conn" xml:"conn"`

	// Departure element. This element does not have child elements. All information about the departure is stored in attributes (see the next table).
	Dp *Event `json:"dp,omitempty" xml:"dp,omitempty"`

	// EVA station number.
	// Required: true
	Eva *int64 `json:"eva" xml:"eva,attr"`

	// Historic delay element.
	Hd []*HistoricDelay `json:"hd" xml:"hd"`

	// Historic platform change element.
	Hpc []*HistoricPlatformChange `json:"hpc" xml:"hpc"`

	// Stop index.
	// Required: true
	I *int64 `json:"i" xml:"i,attr"`

	// Junction type.
	Jt JunctionType `json:"jt,omitempty" xml:"jt,attr,omitempty"`

	// Message element.
	M []*Message `json:"m" xml:"m"`

	// Reference trip relation element.
	Rtr []*ReferenceTripRelation `json:"rtr" xml:"rtr"`
}

// Validate validates this trip stop
func (m *TripStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEva(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TripStop) validateAr(formats strfmt.Registry) error {

	if swag.IsZero(m.Ar) { // not required
		return nil
	}

	if m.Ar != nil {
		if err := m.Ar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ar")
			}
			return err
		}
	}

	return nil
}

func (m *TripStop) validateConn(formats strfmt.Registry) error {

	if swag.IsZero(m.Conn) { // not required
		return nil
	}

	for i := 0; i < len(m.Conn); i++ {
		if swag.IsZero(m.Conn[i]) { // not required
			continue
		}

		if m.Conn[i] != nil {
			if err := m.Conn[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conn" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TripStop) validateDp(formats strfmt.Registry) error {

	if swag.IsZero(m.Dp) { // not required
		return nil
	}

	if m.Dp != nil {
		if err := m.Dp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dp")
			}
			return err
		}
	}

	return nil
}

func (m *TripStop) validateEva(formats strfmt.Registry) error {

	if err := validate.Required("eva", "body", m.Eva); err != nil {
		return err
	}

	return nil
}

func (m *TripStop) validateHd(formats strfmt.Registry) error {

	if swag.IsZero(m.Hd) { // not required
		return nil
	}

	for i := 0; i < len(m.Hd); i++ {
		if swag.IsZero(m.Hd[i]) { // not required
			continue
		}

		if m.Hd[i] != nil {
			if err := m.Hd[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hd" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TripStop) validateHpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Hpc) { // not required
		return nil
	}

	for i := 0; i < len(m.Hpc); i++ {
		if swag.IsZero(m.Hpc[i]) { // not required
			continue
		}

		if m.Hpc[i] != nil {
			if err := m.Hpc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hpc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TripStop) validateI(formats strfmt.Registry) error {

	if err := validate.Required("i", "body", m.I); err != nil {
		return err
	}

	return nil
}

func (m *TripStop) validateJt(formats strfmt.Registry) error {

	if swag.IsZero(m.Jt) { // not required
		return nil
	}

	if err := m.Jt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("jt")
		}
		return err
	}

	return nil
}

func (m *TripStop) validateM(formats strfmt.Registry) error {

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

func (m *TripStop) validateRtr(formats strfmt.Registry) error {

	if swag.IsZero(m.Rtr) { // not required
		return nil
	}

	for i := 0; i < len(m.Rtr); i++ {
		if swag.IsZero(m.Rtr[i]) { // not required
			continue
		}

		if m.Rtr[i] != nil {
			if err := m.Rtr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rtr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TripStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripStop) UnmarshalBinary(b []byte) error {
	var res TripStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}