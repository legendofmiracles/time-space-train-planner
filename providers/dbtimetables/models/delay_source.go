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

// DelaySource Delay source.
//
// * L - LEIBIT
//   LeiBit/LeiDis.
// * NA - RISNE AUT
//   IRIS-NE (automatisch).
// * NM - RISNE MAN
//   IRIS-NE (manuell).
// * V - VDV
//   Prognosen durch dritte EVU über VDVin.
// * IA - ISTP AUT
//   ISTP automatisch.
// * IM - ISTP MAN
//   ISTP manuell.
// * A - AUTOMATIC PROGNOSIS
//   Automatische Prognose durch Prognoseautomat.
//
//
// swagger:model delaySource
type DelaySource string

const (

	// DelaySourceL captures enum value "L"
	DelaySourceL DelaySource = "L"

	// DelaySourceNA captures enum value "NA"
	DelaySourceNA DelaySource = "NA"

	// DelaySourceNM captures enum value "NM"
	DelaySourceNM DelaySource = "NM"

	// DelaySourceV captures enum value "V"
	DelaySourceV DelaySource = "V"

	// DelaySourceIA captures enum value "IA"
	DelaySourceIA DelaySource = "IA"

	// DelaySourceIM captures enum value "IM"
	DelaySourceIM DelaySource = "IM"

	// DelaySourceA captures enum value "A"
	DelaySourceA DelaySource = "A"
)

// for schema
var delaySourceEnum []interface{}

func init() {
	var res []DelaySource
	if err := json.Unmarshal([]byte(`["L","NA","NM","V","IA","IM","A"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		delaySourceEnum = append(delaySourceEnum, v)
	}
}

func (m DelaySource) validateDelaySourceEnum(path, location string, value DelaySource) error {
	if err := validate.Enum(path, location, value, delaySourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this delay source
func (m DelaySource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDelaySourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}