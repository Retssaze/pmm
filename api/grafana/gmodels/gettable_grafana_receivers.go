// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GettableGrafanaReceivers gettable grafana receivers
//
// swagger:model GettableGrafanaReceivers
type GettableGrafanaReceivers struct {

	// grafana managed receivers
	GrafanaManagedReceivers []*GettableGrafanaReceiver `json:"grafana_managed_receiver_configs"`
}

// Validate validates this gettable grafana receivers
func (m *GettableGrafanaReceivers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrafanaManagedReceivers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableGrafanaReceivers) validateGrafanaManagedReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaManagedReceivers) { // not required
		return nil
	}

	for i := 0; i < len(m.GrafanaManagedReceivers); i++ {
		if swag.IsZero(m.GrafanaManagedReceivers[i]) { // not required
			continue
		}

		if m.GrafanaManagedReceivers[i] != nil {
			if err := m.GrafanaManagedReceivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gettable grafana receivers based on the context it is used
func (m *GettableGrafanaReceivers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrafanaManagedReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableGrafanaReceivers) contextValidateGrafanaManagedReceivers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GrafanaManagedReceivers); i++ {

		if m.GrafanaManagedReceivers[i] != nil {
			if err := m.GrafanaManagedReceivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GettableGrafanaReceivers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GettableGrafanaReceivers) UnmarshalBinary(b []byte) error {
	var res GettableGrafanaReceivers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}