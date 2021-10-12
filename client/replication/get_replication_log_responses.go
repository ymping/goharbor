// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ymping/goharbor/models"
)

// GetReplicationLogReader is a Reader for the GetReplicationLog structure.
type GetReplicationLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationLogOK creates a GetReplicationLogOK with default headers values
func NewGetReplicationLogOK() *GetReplicationLogOK {
	return &GetReplicationLogOK{}
}

/* GetReplicationLogOK describes a response with status code 200, with default header values.

Success
*/
type GetReplicationLogOK struct {

	/* The content type of response body
	 */
	ContentType string

	Payload string
}

func (o *GetReplicationLogOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks/{task_id}/log][%d] getReplicationLogOK  %+v", 200, o.Payload)
}
func (o *GetReplicationLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetReplicationLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationLogUnauthorized creates a GetReplicationLogUnauthorized with default headers values
func NewGetReplicationLogUnauthorized() *GetReplicationLogUnauthorized {
	return &GetReplicationLogUnauthorized{}
}

/* GetReplicationLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReplicationLogUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetReplicationLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks/{task_id}/log][%d] getReplicationLogUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReplicationLogUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationLogForbidden creates a GetReplicationLogForbidden with default headers values
func NewGetReplicationLogForbidden() *GetReplicationLogForbidden {
	return &GetReplicationLogForbidden{}
}

/* GetReplicationLogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReplicationLogForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetReplicationLogForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks/{task_id}/log][%d] getReplicationLogForbidden  %+v", 403, o.Payload)
}
func (o *GetReplicationLogForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationLogNotFound creates a GetReplicationLogNotFound with default headers values
func NewGetReplicationLogNotFound() *GetReplicationLogNotFound {
	return &GetReplicationLogNotFound{}
}

/* GetReplicationLogNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetReplicationLogNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetReplicationLogNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks/{task_id}/log][%d] getReplicationLogNotFound  %+v", 404, o.Payload)
}
func (o *GetReplicationLogNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationLogInternalServerError creates a GetReplicationLogInternalServerError with default headers values
func NewGetReplicationLogInternalServerError() *GetReplicationLogInternalServerError {
	return &GetReplicationLogInternalServerError{}
}

/* GetReplicationLogInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetReplicationLogInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetReplicationLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}/tasks/{task_id}/log][%d] getReplicationLogInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReplicationLogInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
