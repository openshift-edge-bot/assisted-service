// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2RegisterClusterReader is a Reader for the V2RegisterCluster structure.
type V2RegisterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2RegisterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2RegisterClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2RegisterClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2RegisterClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2RegisterClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2RegisterClusterMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2RegisterClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2RegisterClusterCreated creates a V2RegisterClusterCreated with default headers values
func NewV2RegisterClusterCreated() *V2RegisterClusterCreated {
	return &V2RegisterClusterCreated{}
}

/*V2RegisterClusterCreated handles this case with default header values.

Success.
*/
type V2RegisterClusterCreated struct {
	Payload *models.Cluster
}

func (o *V2RegisterClusterCreated) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterCreated  %+v", 201, o.Payload)
}

func (o *V2RegisterClusterCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *V2RegisterClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2RegisterClusterBadRequest creates a V2RegisterClusterBadRequest with default headers values
func NewV2RegisterClusterBadRequest() *V2RegisterClusterBadRequest {
	return &V2RegisterClusterBadRequest{}
}

/*V2RegisterClusterBadRequest handles this case with default header values.

Error.
*/
type V2RegisterClusterBadRequest struct {
	Payload *models.Error
}

func (o *V2RegisterClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterBadRequest  %+v", 400, o.Payload)
}

func (o *V2RegisterClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2RegisterClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2RegisterClusterUnauthorized creates a V2RegisterClusterUnauthorized with default headers values
func NewV2RegisterClusterUnauthorized() *V2RegisterClusterUnauthorized {
	return &V2RegisterClusterUnauthorized{}
}

/*V2RegisterClusterUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2RegisterClusterUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2RegisterClusterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *V2RegisterClusterUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2RegisterClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2RegisterClusterForbidden creates a V2RegisterClusterForbidden with default headers values
func NewV2RegisterClusterForbidden() *V2RegisterClusterForbidden {
	return &V2RegisterClusterForbidden{}
}

/*V2RegisterClusterForbidden handles this case with default header values.

Forbidden.
*/
type V2RegisterClusterForbidden struct {
	Payload *models.InfraError
}

func (o *V2RegisterClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterForbidden  %+v", 403, o.Payload)
}

func (o *V2RegisterClusterForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2RegisterClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2RegisterClusterMethodNotAllowed creates a V2RegisterClusterMethodNotAllowed with default headers values
func NewV2RegisterClusterMethodNotAllowed() *V2RegisterClusterMethodNotAllowed {
	return &V2RegisterClusterMethodNotAllowed{}
}

/*V2RegisterClusterMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2RegisterClusterMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2RegisterClusterMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2RegisterClusterMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2RegisterClusterMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2RegisterClusterInternalServerError creates a V2RegisterClusterInternalServerError with default headers values
func NewV2RegisterClusterInternalServerError() *V2RegisterClusterInternalServerError {
	return &V2RegisterClusterInternalServerError{}
}

/*V2RegisterClusterInternalServerError handles this case with default header values.

Error.
*/
type V2RegisterClusterInternalServerError struct {
	Payload *models.Error
}

func (o *V2RegisterClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters][%d] v2RegisterClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *V2RegisterClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2RegisterClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
