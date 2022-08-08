// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// GetPresignedForAssistedServiceISOReader is a Reader for the GetPresignedForAssistedServiceISO structure.
type GetPresignedForAssistedServiceISOReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPresignedForAssistedServiceISOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPresignedForAssistedServiceISOOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPresignedForAssistedServiceISOUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPresignedForAssistedServiceISOForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPresignedForAssistedServiceISONotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPresignedForAssistedServiceISOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPresignedForAssistedServiceISOOK creates a GetPresignedForAssistedServiceISOOK with default headers values
func NewGetPresignedForAssistedServiceISOOK() *GetPresignedForAssistedServiceISOOK {
	return &GetPresignedForAssistedServiceISOOK{}
}

/* GetPresignedForAssistedServiceISOOK describes a response with status code 200, with default header values.

Success.
*/
type GetPresignedForAssistedServiceISOOK struct {
	Payload *models.Presigned
}

func (o *GetPresignedForAssistedServiceISOOK) Error() string {
	return fmt.Sprintf("[GET /v1/assisted-service-iso/presigned][%d] getPresignedForAssistedServiceISOOK  %+v", 200, o.Payload)
}
func (o *GetPresignedForAssistedServiceISOOK) GetPayload() *models.Presigned {
	return o.Payload
}

func (o *GetPresignedForAssistedServiceISOOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Presigned)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresignedForAssistedServiceISOUnauthorized creates a GetPresignedForAssistedServiceISOUnauthorized with default headers values
func NewGetPresignedForAssistedServiceISOUnauthorized() *GetPresignedForAssistedServiceISOUnauthorized {
	return &GetPresignedForAssistedServiceISOUnauthorized{}
}

/* GetPresignedForAssistedServiceISOUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetPresignedForAssistedServiceISOUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetPresignedForAssistedServiceISOUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/assisted-service-iso/presigned][%d] getPresignedForAssistedServiceISOUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPresignedForAssistedServiceISOUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetPresignedForAssistedServiceISOUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresignedForAssistedServiceISOForbidden creates a GetPresignedForAssistedServiceISOForbidden with default headers values
func NewGetPresignedForAssistedServiceISOForbidden() *GetPresignedForAssistedServiceISOForbidden {
	return &GetPresignedForAssistedServiceISOForbidden{}
}

/* GetPresignedForAssistedServiceISOForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type GetPresignedForAssistedServiceISOForbidden struct {
	Payload *models.InfraError
}

func (o *GetPresignedForAssistedServiceISOForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/assisted-service-iso/presigned][%d] getPresignedForAssistedServiceISOForbidden  %+v", 403, o.Payload)
}
func (o *GetPresignedForAssistedServiceISOForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetPresignedForAssistedServiceISOForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresignedForAssistedServiceISONotFound creates a GetPresignedForAssistedServiceISONotFound with default headers values
func NewGetPresignedForAssistedServiceISONotFound() *GetPresignedForAssistedServiceISONotFound {
	return &GetPresignedForAssistedServiceISONotFound{}
}

/* GetPresignedForAssistedServiceISONotFound describes a response with status code 404, with default header values.

Error.
*/
type GetPresignedForAssistedServiceISONotFound struct {
	Payload *models.Error
}

func (o *GetPresignedForAssistedServiceISONotFound) Error() string {
	return fmt.Sprintf("[GET /v1/assisted-service-iso/presigned][%d] getPresignedForAssistedServiceISONotFound  %+v", 404, o.Payload)
}
func (o *GetPresignedForAssistedServiceISONotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPresignedForAssistedServiceISONotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresignedForAssistedServiceISOInternalServerError creates a GetPresignedForAssistedServiceISOInternalServerError with default headers values
func NewGetPresignedForAssistedServiceISOInternalServerError() *GetPresignedForAssistedServiceISOInternalServerError {
	return &GetPresignedForAssistedServiceISOInternalServerError{}
}

/* GetPresignedForAssistedServiceISOInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type GetPresignedForAssistedServiceISOInternalServerError struct {
	Payload *models.Error
}

func (o *GetPresignedForAssistedServiceISOInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/assisted-service-iso/presigned][%d] getPresignedForAssistedServiceISOInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPresignedForAssistedServiceISOInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPresignedForAssistedServiceISOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}