// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertNotificationLookup alert notification lookup
//
// swagger:model AlertNotificationLookup
type AlertNotificationLookup struct {

	// Id
	ID int64 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// Uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this alert notification lookup
func (m *AlertNotificationLookup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alert notification lookup based on context it is used
func (m *AlertNotificationLookup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertNotificationLookup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertNotificationLookup) UnmarshalBinary(b []byte) error {
	var res AlertNotificationLookup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}