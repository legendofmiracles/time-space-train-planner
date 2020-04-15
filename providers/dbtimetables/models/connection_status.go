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

// ConnectionStatus Connection status.
//
// * w - WAITING
//   This (regular) connection is waiting.
// * n - TRANSITION
//   This (regular) connection CANNOT wait.
// * a - ALTERNATIVE
//   This is an alternative (unplanned) connection that has been introduced as a replacement for one regular connection that cannot wait. The connections "tl" (triplabel) attribute might in this case refer to the replaced connection (or more specifi-cally the trip from that connection). Alternative connections are always waiting (they are re-moved otherwise).
//
//
// swagger:model connectionStatus
type ConnectionStatus string

const (

	// ConnectionStatusW captures enum value "w"
	ConnectionStatusW ConnectionStatus = "w"

	// ConnectionStatusN captures enum value "n"
	ConnectionStatusN ConnectionStatus = "n"

	// ConnectionStatusA captures enum value "a"
	ConnectionStatusA ConnectionStatus = "a"
)

// for schema
var connectionStatusEnum []interface{}

func init() {
	var res []ConnectionStatus
	if err := json.Unmarshal([]byte(`["w","n","a"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionStatusEnum = append(connectionStatusEnum, v)
	}
}

func (m ConnectionStatus) validateConnectionStatusEnum(path, location string, value ConnectionStatus) error {
	if err := validate.Enum(path, location, value, connectionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this connection status
func (m ConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConnectionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}