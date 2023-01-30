// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadHostIgnitionOKCode is the HTTP code returned for type DownloadHostIgnitionOK
const DownloadHostIgnitionOKCode int = 200

/*DownloadHostIgnitionOK Success.

swagger:response downloadHostIgnitionOK
*/
type DownloadHostIgnitionOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadHostIgnitionOK creates DownloadHostIgnitionOK with default headers values
func NewDownloadHostIgnitionOK() *DownloadHostIgnitionOK {

	return &DownloadHostIgnitionOK{}
}

// WithPayload adds the payload to the download host ignition o k response
func (o *DownloadHostIgnitionOK) WithPayload(payload io.ReadCloser) *DownloadHostIgnitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition o k response
func (o *DownloadHostIgnitionOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadHostIgnitionUnauthorizedCode is the HTTP code returned for type DownloadHostIgnitionUnauthorized
const DownloadHostIgnitionUnauthorizedCode int = 401

/*DownloadHostIgnitionUnauthorized Unauthorized.

swagger:response downloadHostIgnitionUnauthorized
*/
type DownloadHostIgnitionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadHostIgnitionUnauthorized creates DownloadHostIgnitionUnauthorized with default headers values
func NewDownloadHostIgnitionUnauthorized() *DownloadHostIgnitionUnauthorized {

	return &DownloadHostIgnitionUnauthorized{}
}

// WithPayload adds the payload to the download host ignition unauthorized response
func (o *DownloadHostIgnitionUnauthorized) WithPayload(payload *models.InfraError) *DownloadHostIgnitionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition unauthorized response
func (o *DownloadHostIgnitionUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionForbiddenCode is the HTTP code returned for type DownloadHostIgnitionForbidden
const DownloadHostIgnitionForbiddenCode int = 403

/*DownloadHostIgnitionForbidden Forbidden.

swagger:response downloadHostIgnitionForbidden
*/
type DownloadHostIgnitionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadHostIgnitionForbidden creates DownloadHostIgnitionForbidden with default headers values
func NewDownloadHostIgnitionForbidden() *DownloadHostIgnitionForbidden {

	return &DownloadHostIgnitionForbidden{}
}

// WithPayload adds the payload to the download host ignition forbidden response
func (o *DownloadHostIgnitionForbidden) WithPayload(payload *models.InfraError) *DownloadHostIgnitionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition forbidden response
func (o *DownloadHostIgnitionForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionNotFoundCode is the HTTP code returned for type DownloadHostIgnitionNotFound
const DownloadHostIgnitionNotFoundCode int = 404

/*DownloadHostIgnitionNotFound Error.

swagger:response downloadHostIgnitionNotFound
*/
type DownloadHostIgnitionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadHostIgnitionNotFound creates DownloadHostIgnitionNotFound with default headers values
func NewDownloadHostIgnitionNotFound() *DownloadHostIgnitionNotFound {

	return &DownloadHostIgnitionNotFound{}
}

// WithPayload adds the payload to the download host ignition not found response
func (o *DownloadHostIgnitionNotFound) WithPayload(payload *models.Error) *DownloadHostIgnitionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition not found response
func (o *DownloadHostIgnitionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionMethodNotAllowedCode is the HTTP code returned for type DownloadHostIgnitionMethodNotAllowed
const DownloadHostIgnitionMethodNotAllowedCode int = 405

/*DownloadHostIgnitionMethodNotAllowed Method Not Allowed.

swagger:response downloadHostIgnitionMethodNotAllowed
*/
type DownloadHostIgnitionMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadHostIgnitionMethodNotAllowed creates DownloadHostIgnitionMethodNotAllowed with default headers values
func NewDownloadHostIgnitionMethodNotAllowed() *DownloadHostIgnitionMethodNotAllowed {

	return &DownloadHostIgnitionMethodNotAllowed{}
}

// WithPayload adds the payload to the download host ignition method not allowed response
func (o *DownloadHostIgnitionMethodNotAllowed) WithPayload(payload *models.Error) *DownloadHostIgnitionMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition method not allowed response
func (o *DownloadHostIgnitionMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionConflictCode is the HTTP code returned for type DownloadHostIgnitionConflict
const DownloadHostIgnitionConflictCode int = 409

/*DownloadHostIgnitionConflict Error.

swagger:response downloadHostIgnitionConflict
*/
type DownloadHostIgnitionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadHostIgnitionConflict creates DownloadHostIgnitionConflict with default headers values
func NewDownloadHostIgnitionConflict() *DownloadHostIgnitionConflict {

	return &DownloadHostIgnitionConflict{}
}

// WithPayload adds the payload to the download host ignition conflict response
func (o *DownloadHostIgnitionConflict) WithPayload(payload *models.Error) *DownloadHostIgnitionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition conflict response
func (o *DownloadHostIgnitionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionInternalServerErrorCode is the HTTP code returned for type DownloadHostIgnitionInternalServerError
const DownloadHostIgnitionInternalServerErrorCode int = 500

/*DownloadHostIgnitionInternalServerError Error.

swagger:response downloadHostIgnitionInternalServerError
*/
type DownloadHostIgnitionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadHostIgnitionInternalServerError creates DownloadHostIgnitionInternalServerError with default headers values
func NewDownloadHostIgnitionInternalServerError() *DownloadHostIgnitionInternalServerError {

	return &DownloadHostIgnitionInternalServerError{}
}

// WithPayload adds the payload to the download host ignition internal server error response
func (o *DownloadHostIgnitionInternalServerError) WithPayload(payload *models.Error) *DownloadHostIgnitionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition internal server error response
func (o *DownloadHostIgnitionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadHostIgnitionServiceUnavailableCode is the HTTP code returned for type DownloadHostIgnitionServiceUnavailable
const DownloadHostIgnitionServiceUnavailableCode int = 503

/*DownloadHostIgnitionServiceUnavailable Unavailable.

swagger:response downloadHostIgnitionServiceUnavailable
*/
type DownloadHostIgnitionServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadHostIgnitionServiceUnavailable creates DownloadHostIgnitionServiceUnavailable with default headers values
func NewDownloadHostIgnitionServiceUnavailable() *DownloadHostIgnitionServiceUnavailable {

	return &DownloadHostIgnitionServiceUnavailable{}
}

// WithPayload adds the payload to the download host ignition service unavailable response
func (o *DownloadHostIgnitionServiceUnavailable) WithPayload(payload *models.Error) *DownloadHostIgnitionServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download host ignition service unavailable response
func (o *DownloadHostIgnitionServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadHostIgnitionServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}