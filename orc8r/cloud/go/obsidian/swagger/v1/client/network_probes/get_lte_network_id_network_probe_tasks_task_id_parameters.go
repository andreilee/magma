// Code generated by go-swagger; DO NOT EDIT.

package network_probes

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

// NewGetLTENetworkIDNetworkProbeTasksTaskIDParams creates a new GetLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized.
func NewGetLTENetworkIDNetworkProbeTasksTaskIDParams() *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeTasksTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithTimeout creates a new GetLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeTasksTaskIDParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithContext creates a new GetLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeTasksTaskIDParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithHTTPClient creates a new GetLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDNetworkProbeTasksTaskIDParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDNetworkProbeTasksTaskIDParams contains all the parameters to send to the API endpoint
for the get LTE network ID network probe tasks task ID operation typically these are written to a http.Request
*/
type GetLTENetworkIDNetworkProbeTasksTaskIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TaskID
	  Network Probe Task ID

	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WithNetworkID(networkID string) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTaskID adds the taskID to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WithTaskID(taskID string) *GetLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get LTE network ID network probe tasks task ID params
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDNetworkProbeTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
