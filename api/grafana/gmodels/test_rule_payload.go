// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestRulePayload test rule payload
//
// swagger:model TestRulePayload
type TestRulePayload struct {

	// expr
	// Example: (node_filesystem_avail_bytes{fstype!=\"\",job=\"integrations/node_exporter\"} node_filesystem_size_bytes{fstype!=\"\",job=\"integrations/node_exporter\"} * 100 \u003c 5 and node_filesystem_readonly{fstype!=\"\",job=\"integrations/node_exporter\"} == 0)
	Expr string `json:"expr,omitempty"`

	// grafana condition
	GrafanaCondition *EvalAlertConditionCommand `json:"grafana_condition,omitempty"`
}

// Validate validates this test rule payload
func (m *TestRulePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrafanaCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRulePayload) validateGrafanaCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaCondition) { // not required
		return nil
	}

	if m.GrafanaCondition != nil {
		if err := m.GrafanaCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_condition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test rule payload based on the context it is used
func (m *TestRulePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrafanaCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestRulePayload) contextValidateGrafanaCondition(ctx context.Context, formats strfmt.Registry) error {

	if m.GrafanaCondition != nil {
		if err := m.GrafanaCondition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_condition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestRulePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRulePayload) UnmarshalBinary(b []byte) error {
	var res TestRulePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}