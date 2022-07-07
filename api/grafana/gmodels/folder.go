// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Folder folder
//
// swagger:model Folder
type Folder struct {

	// can admin
	CanAdmin bool `json:"canAdmin,omitempty"`

	// can delete
	CanDelete bool `json:"canDelete,omitempty"`

	// can edit
	CanEdit bool `json:"canEdit,omitempty"`

	// can save
	CanSave bool `json:"canSave,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// has Acl
	HasACL bool `json:"hasAcl,omitempty"`

	// Id
	ID int64 `json:"id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// Uid
	UID string `json:"uid,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Url
	URL string `json:"url,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this folder
func (m *Folder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Folder) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Folder) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this folder based on context it is used
func (m *Folder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Folder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Folder) UnmarshalBinary(b []byte) error {
	var res Folder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}