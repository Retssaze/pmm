// Code generated by go-swagger; DO NOT EDIT.

package server

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
)

// NewGetFileParams creates a new GetFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileParams() *GetFileParams {
	return &GetFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileParamsWithTimeout creates a new GetFileParams object
// with the ability to set a timeout on a request.
func NewGetFileParamsWithTimeout(timeout time.Duration) *GetFileParams {
	return &GetFileParams{
		timeout: timeout,
	}
}

// NewGetFileParamsWithContext creates a new GetFileParams object
// with the ability to set a context for a request.
func NewGetFileParamsWithContext(ctx context.Context) *GetFileParams {
	return &GetFileParams{
		Context: ctx,
	}
}

// NewGetFileParamsWithHTTPClient creates a new GetFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileParamsWithHTTPClient(client *http.Client) *GetFileParams {
	return &GetFileParams{
		HTTPClient: client,
	}
}

/*
GetFileParams contains all the parameters to send to the API endpoint

	for the get file operation.

	Typically these are written to a http.Request.
*/
type GetFileParams struct {
	// Body.
	Body GetFileBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileParams) WithDefaults() *GetFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file params
func (o *GetFileParams) WithTimeout(timeout time.Duration) *GetFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file params
func (o *GetFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file params
func (o *GetFileParams) WithContext(ctx context.Context) *GetFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file params
func (o *GetFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file params
func (o *GetFileParams) WithHTTPClient(client *http.Client) *GetFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file params
func (o *GetFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get file params
func (o *GetFileParams) WithBody(body GetFileBody) *GetFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get file params
func (o *GetFileParams) SetBody(body GetFileBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}