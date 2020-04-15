// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EventStatus Event status.
//
// * p - PLANNED
//   The event was planned. This status is also used when the cancellation of an event has been revoked.
// * a - ADDED
//   The event was added to the planned data (new stop).
// * c - CANCELLED
//   The event was canceled (as changedstatus, can apply to planned and added stops).
//
//
// swagger:model eventStatus
type EventStatus string

const (

	// EventStatusP captures enum value "p"
	EventStatusP EventStatus = "p"

	// EventStatusA captures enum value "a"
	EventStatusA EventStatus = "a"

	// EventStatusC captures enum value "c"
	EventStatusC EventStatus = "c"
)

// for schema
var eventStatusEnum []interface{}

func init() {
	var res []EventStatus
	if err := json.Unmarshal([]byte(`["p","a","c"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventStatusEnum = append(eventStatusEnum, v)
	}
}

func (m EventStatus) validateEventStatusEnum(path, location string, value EventStatus) error {
	if err := validate.Enum(path, location, value, eventStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this event status
func (m EventStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEventStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}