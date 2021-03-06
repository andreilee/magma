// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWifiParams creates a new GetWifiParams object
// with the default values initialized.
func NewGetWifiParams() *GetWifiParams {

	return &GetWifiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWifiParamsWithTimeout creates a new GetWifiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWifiParamsWithTimeout(timeout time.Duration) *GetWifiParams {

	return &GetWifiParams{

		timeout: timeout,
	}
}

// NewGetWifiParamsWithContext creates a new GetWifiParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWifiParamsWithContext(ctx context.Context) *GetWifiParams {

	return &GetWifiParams{

		Context: ctx,
	}
}

// NewGetWifiParamsWithHTTPClient creates a new GetWifiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWifiParamsWithHTTPClient(client *http.Client) *GetWifiParams {

	return &GetWifiParams{
		HTTPClient: client,
	}
}

/*GetWifiParams contains all the parameters to send to the API endpoint
for the get wifi operation typically these are written to a http.Request
*/
type GetWifiParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wifi params
func (o *GetWifiParams) WithTimeout(timeout time.Duration) *GetWifiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wifi params
func (o *GetWifiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wifi params
func (o *GetWifiParams) WithContext(ctx context.Context) *GetWifiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wifi params
func (o *GetWifiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wifi params
func (o *GetWifiParams) WithHTTPClient(client *http.Client) *GetWifiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wifi params
func (o *GetWifiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWifiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
