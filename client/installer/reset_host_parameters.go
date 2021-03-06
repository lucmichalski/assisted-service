// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewResetHostParams creates a new ResetHostParams object
// with the default values initialized.
func NewResetHostParams() *ResetHostParams {
	var ()
	return &ResetHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetHostParamsWithTimeout creates a new ResetHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetHostParamsWithTimeout(timeout time.Duration) *ResetHostParams {
	var ()
	return &ResetHostParams{

		timeout: timeout,
	}
}

// NewResetHostParamsWithContext creates a new ResetHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewResetHostParamsWithContext(ctx context.Context) *ResetHostParams {
	var ()
	return &ResetHostParams{

		Context: ctx,
	}
}

// NewResetHostParamsWithHTTPClient creates a new ResetHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetHostParamsWithHTTPClient(client *http.Client) *ResetHostParams {
	var ()
	return &ResetHostParams{
		HTTPClient: client,
	}
}

/*ResetHostParams contains all the parameters to send to the API endpoint
for the reset host operation typically these are written to a http.Request
*/
type ResetHostParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID
	/*HostID*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset host params
func (o *ResetHostParams) WithTimeout(timeout time.Duration) *ResetHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset host params
func (o *ResetHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset host params
func (o *ResetHostParams) WithContext(ctx context.Context) *ResetHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset host params
func (o *ResetHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset host params
func (o *ResetHostParams) WithHTTPClient(client *http.Client) *ResetHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset host params
func (o *ResetHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the reset host params
func (o *ResetHostParams) WithClusterID(clusterID strfmt.UUID) *ResetHostParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the reset host params
func (o *ResetHostParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the reset host params
func (o *ResetHostParams) WithHostID(hostID strfmt.UUID) *ResetHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the reset host params
func (o *ResetHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
