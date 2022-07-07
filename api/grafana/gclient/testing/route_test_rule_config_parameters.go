// Code generated by go-swagger; DO NOT EDIT.

package testing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// NewRouteTestRuleConfigParams creates a new RouteTestRuleConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteTestRuleConfigParams() *RouteTestRuleConfigParams {
	return &RouteTestRuleConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteTestRuleConfigParamsWithTimeout creates a new RouteTestRuleConfigParams object
// with the ability to set a timeout on a request.
func NewRouteTestRuleConfigParamsWithTimeout(timeout time.Duration) *RouteTestRuleConfigParams {
	return &RouteTestRuleConfigParams{
		timeout: timeout,
	}
}

// NewRouteTestRuleConfigParamsWithContext creates a new RouteTestRuleConfigParams object
// with the ability to set a context for a request.
func NewRouteTestRuleConfigParamsWithContext(ctx context.Context) *RouteTestRuleConfigParams {
	return &RouteTestRuleConfigParams{
		Context: ctx,
	}
}

// NewRouteTestRuleConfigParamsWithHTTPClient creates a new RouteTestRuleConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteTestRuleConfigParamsWithHTTPClient(client *http.Client) *RouteTestRuleConfigParams {
	return &RouteTestRuleConfigParams{
		HTTPClient: client,
	}
}

/* RouteTestRuleConfigParams contains all the parameters to send to the API endpoint
   for the route test rule config operation.

   Typically these are written to a http.Request.
*/
type RouteTestRuleConfigParams struct {

	// Body.
	Body *gmodels.TestRulePayload

	/* Recipient.

	     Recipient should be "grafana" for requests to be handled by grafana
	and the numeric datasource id for requests to be forwarded to a datasource
	*/
	Recipient string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route test rule config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteTestRuleConfigParams) WithDefaults() *RouteTestRuleConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route test rule config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteTestRuleConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route test rule config params
func (o *RouteTestRuleConfigParams) WithTimeout(timeout time.Duration) *RouteTestRuleConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route test rule config params
func (o *RouteTestRuleConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route test rule config params
func (o *RouteTestRuleConfigParams) WithContext(ctx context.Context) *RouteTestRuleConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route test rule config params
func (o *RouteTestRuleConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route test rule config params
func (o *RouteTestRuleConfigParams) WithHTTPClient(client *http.Client) *RouteTestRuleConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route test rule config params
func (o *RouteTestRuleConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the route test rule config params
func (o *RouteTestRuleConfigParams) WithBody(body *gmodels.TestRulePayload) *RouteTestRuleConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the route test rule config params
func (o *RouteTestRuleConfigParams) SetBody(body *gmodels.TestRulePayload) {
	o.Body = body
}

// WithRecipient adds the recipient to the route test rule config params
func (o *RouteTestRuleConfigParams) WithRecipient(recipient string) *RouteTestRuleConfigParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the route test rule config params
func (o *RouteTestRuleConfigParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WriteToRequest writes these params to a swagger request
func (o *RouteTestRuleConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param Recipient
	if err := r.SetPathParam("Recipient", o.Recipient); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}