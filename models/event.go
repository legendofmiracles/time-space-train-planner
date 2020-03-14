// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Event An event (arrival or departure) that is part of a stop.
//
// swagger:model event
type Event struct {

	// Changed distant endpoint.
	Cde string `json:"cde,omitempty" xml:"cde,attr,omitempty"`

	// Cancellation time. Time when the cancellation of this stop was created. The time, in ten digit 'YYMMddHHmm' format, e.g. '1404011437' for 14:37 on April the 1st of 2014.
	Clt string `json:"clt,omitempty" xml:"clt,attr,omitempty"`

	// Changed platform.
	Cp string `json:"cp,omitempty" xml:"cp,attr,omitempty"`

	// Changed path.
	Cpth string `json:"cpth,omitempty" xml:"cpth,attr,omitempty"`

	// Changed status. The status of this event, a one-character indicator that is one of:
	// * 'a' = this event was added
	// * 'c' = this event was cancelled
	// * 'p' = this event was planned (also used when the cancellation of an event has been revoked)
	// The status applies to the event, not to the trip as a whole. Insertion or removal of a single stop will usually affect two events at once: one arrival and one departure event. Note that these two events do not have to belong to the same stop. For example, removing the last stop of a trip will result in arrival cancellation for the last stop and of departure cancellation for the stop before the last. So asymmetric cancellations of just arrival or departure for a stop can occur.
	//
	Cs EventStatus `json:"cs,omitempty" xml:"cs,attr,omitempty"`

	// Changed time. New estimated or actual departure or arrival time. The time, in ten digit 'YYMMddHHmm' format, e.g. '1404011437' for 14:37 on April the 1st of 2014.
	Ct string `json:"ct,omitempty" xml:"ct,attr,omitempty"`

	// Distant change.
	Dc int64 `json:"dc,omitempty" xml:"dc,attr,omitempty"`

	// Hidden. 1 if the event should not be shown on WBT because travellers are not supposed to enter or exit the train at this stop.
	Hi int64 `json:"hi,omitempty" xml:"hi,attr,omitempty"`

	// Line. The line indicator (e.g. "3" for an S-Bahn or "45S" for a bus).
	L string `json:"l,omitempty" xml:"l,attr,omitempty"`

	// List of messages.
	M []*Message `json:"m" xml:"m"`

	// Planned distant endpoint.
	Pde string `json:"pde,omitempty" xml:"pde,attr,omitempty"`

	// Planned platform.
	Pp string `json:"pp,omitempty" xml:"pp,attr,omitempty"`

	// Planned Path. A sequence of station names separated by the pipe symbols ('|').
	// E.g.: 'Mainz Hbf|R�sselsheim|Frankfrt(M) Flughafen'.
	// For arrival, the path indicates the stations that come before the current station. The first element then is the trip's start station.
	// For departure, the path indicates the stations that come after the current station. The last element in the path then is the trip's destination station.
	// Note that the current station is never included in the path (neither for arrival nor for departure).
	//
	Ppth string `json:"ppth,omitempty" xml:"ppth,attr,omitempty"`

	// Planned status.
	Ps EventStatus `json:"ps,omitempty" xml:"ps,attr,omitempty"`

	// Planned time. Planned departure or arrival time. The time, in ten digit 'YYMMddHHmm' format, e.g. '1404011437' for 14:37 on April the 1st of 2014.
	Pt string `json:"pt,omitempty" xml:"pt,attr,omitempty"`

	// Transition. Trip id of the next or previous train of a shared train. At the start stop this references the previous trip, at the last stop it references the next trip. E.g. '2016448009055686515-1403311438-1'
	Tra string `json:"tra,omitempty" xml:"tra,omitempty"`

	// Wing. A sequence of trip id separated by the pipe symbols ('|'). E.g. '-906407760000782942-1403311431'.
	Wings string `json:"wings,omitempty" xml:"wings,attr,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCs(formats strfmt.Registry) error {

	if swag.IsZero(m.Cs) { // not required
		return nil
	}

	if err := m.Cs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cs")
		}
		return err
	}

	return nil
}

func (m *Event) validateM(formats strfmt.Registry) error {

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

func (m *Event) validatePs(formats strfmt.Registry) error {

	if swag.IsZero(m.Ps) { // not required
		return nil
	}

	if err := m.Ps.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ps")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
