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

// NewGetClusterInstallConfigParams creates a new GetClusterInstallConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterInstallConfigParams() *GetClusterInstallConfigParams {
	return &GetClusterInstallConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterInstallConfigParamsWithTimeout creates a new GetClusterInstallConfigParams object
// with the ability to set a timeout on a request.
func NewGetClusterInstallConfigParamsWithTimeout(timeout time.Duration) *GetClusterInstallConfigParams {
	return &GetClusterInstallConfigParams{
		timeout: timeout,
	}
}

// NewGetClusterInstallConfigParamsWithContext creates a new GetClusterInstallConfigParams object
// with the ability to set a context for a request.
func NewGetClusterInstallConfigParamsWithContext(ctx context.Context) *GetClusterInstallConfigParams {
	return &GetClusterInstallConfigParams{
		Context: ctx,
	}
}

// NewGetClusterInstallConfigParamsWithHTTPClient creates a new GetClusterInstallConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterInstallConfigParamsWithHTTPClient(client *http.Client) *GetClusterInstallConfigParams {
	return &GetClusterInstallConfigParams{
		HTTPClient: client,
	}
}

/* GetClusterInstallConfigParams contains all the parameters to send to the API endpoint
   for the get cluster install config operation.

   Typically these are written to a http.Request.
*/
type GetClusterInstallConfigParams struct {

	/* ClusterID.

	   The cluster whose install config is being retrieved.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster install config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterInstallConfigParams) WithDefaults() *GetClusterInstallConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster install config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterInstallConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster install config params
func (o *GetClusterInstallConfigParams) WithTimeout(timeout time.Duration) *GetClusterInstallConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster install config params
func (o *GetClusterInstallConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster install config params
func (o *GetClusterInstallConfigParams) WithContext(ctx context.Context) *GetClusterInstallConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster install config params
func (o *GetClusterInstallConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster install config params
func (o *GetClusterInstallConfigParams) WithHTTPClient(client *http.Client) *GetClusterInstallConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster install config params
func (o *GetClusterInstallConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster install config params
func (o *GetClusterInstallConfigParams) WithClusterID(clusterID strfmt.UUID) *GetClusterInstallConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster install config params
func (o *GetClusterInstallConfigParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterInstallConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}