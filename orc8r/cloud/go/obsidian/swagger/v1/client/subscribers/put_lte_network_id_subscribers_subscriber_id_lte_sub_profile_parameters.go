// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams creates a new PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams object
// with the default values initialized.
func NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams() *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	var ()
	return &PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithTimeout creates a new PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	var ()
	return &PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithContext creates a new PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithContext(ctx context.Context) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	var ()
	return &PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithHTTPClient creates a new PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDSubscribersSubscriberIDLTESubProfileParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	var ()
	return &PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams contains all the parameters to send to the API endpoint
for the put LTE network ID subscribers subscriber ID LTE sub profile operation typically these are written to a http.Request
*/
type PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*ProfileName
	  New profile name

	*/
	ProfileName models.SubProfile
	/*SubscriberID
	  Subscriber ID

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithContext(ctx context.Context) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithNetworkID(networkID string) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithProfileName adds the profileName to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithProfileName(profileName models.SubProfile) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetProfileName(profileName)
	return o
}

// SetProfileName adds the profileName to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetProfileName(profileName models.SubProfile) {
	o.ProfileName = profileName
}

// WithSubscriberID adds the subscriberID to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WithSubscriberID(subscriberID string) *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the put LTE network ID subscribers subscriber ID LTE sub profile params
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDSubscribersSubscriberIDLTESubProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ProfileName); err != nil {
		return err
	}

	// path param subscriber_id
	if err := r.SetPathParam("subscriber_id", o.SubscriberID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
