// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error error
//
// swagger:model Error
type Error struct {

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// error status code
	// Required: true
	ErrorStatusCode *int64 `json:"error_status_code"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateErrorStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("error_status_code", "body", m.ErrorStatusCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error based on context it is used
func (m *Error) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
