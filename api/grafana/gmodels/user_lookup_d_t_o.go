// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserLookupDTO user lookup d t o
//
// swagger:model UserLookupDTO
type UserLookupDTO struct {

	// avatar URL
	AvatarURL string `json:"avatarUrl,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// user ID
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this user lookup d t o
func (m *UserLookupDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user lookup d t o based on context it is used
func (m *UserLookupDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserLookupDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLookupDTO) UnmarshalBinary(b []byte) error {
	var res UserLookupDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}