// Code generated by go-swagger; DO NOT EDIT.

package events

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

	"github.com/openshift/assisted-service/models"
)

// NewV2AddEventParams creates a new V2AddEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2AddEventParams() *V2AddEventParams {
	return &V2AddEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2AddEventParamsWithTimeout creates a new V2AddEventParams object
// with the ability to set a timeout on a request.
func NewV2AddEventParamsWithTimeout(timeout time.Duration) *V2AddEventParams {
	return &V2AddEventParams{
		timeout: timeout,
	}
}

// NewV2AddEventParamsWithContext creates a new V2AddEventParams object
// with the ability to set a context for a request.
func NewV2AddEventParamsWithContext(ctx context.Context) *V2AddEventParams {
	return &V2AddEventParams{
		Context: ctx,
	}
}

// NewV2AddEventParamsWithHTTPClient creates a new V2AddEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2AddEventParamsWithHTTPClient(client *http.Client) *V2AddEventParams {
	return &V2AddEventParams{
		HTTPClient: client,
	}
}

/*
V2AddEventParams contains all the parameters to send to the API endpoint

	for the v2 add event operation.

	Typically these are written to a http.Request.
*/
type V2AddEventParams struct {

	/* NewEventParams.

	   The event to be created.
	*/
	NewEventParams *models.Event

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 add event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2AddEventParams) WithDefaults() *V2AddEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 add event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2AddEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 add event params
func (o *V2AddEventParams) WithTimeout(timeout time.Duration) *V2AddEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 add event params
func (o *V2AddEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 add event params
func (o *V2AddEventParams) WithContext(ctx context.Context) *V2AddEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 add event params
func (o *V2AddEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 add event params
func (o *V2AddEventParams) WithHTTPClient(client *http.Client) *V2AddEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 add event params
func (o *V2AddEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewEventParams adds the newEventParams to the v2 add event params
func (o *V2AddEventParams) WithNewEventParams(newEventParams *models.Event) *V2AddEventParams {
	o.SetNewEventParams(newEventParams)
	return o
}

// SetNewEventParams adds the newEventParams to the v2 add event params
func (o *V2AddEventParams) SetNewEventParams(newEventParams *models.Event) {
	o.NewEventParams = newEventParams
}

// WriteToRequest writes these params to a swagger request
func (o *V2AddEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewEventParams != nil {
		if err := r.SetBodyParam(o.NewEventParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}