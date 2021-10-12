// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ymping/goharbor/models"
)

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {
	return &GetLogsOK{}
}

/* GetLogsOK describes a response with status code 200, with default header values.

Success
*/
type GetLogsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of auditlogs
	 */
	XTotalCount int64

	Payload []*models.AuditLog
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsOK  %+v", 200, o.Payload)
}
func (o *GetLogsOK) GetPayload() []*models.AuditLog {
	return o.Payload
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsBadRequest creates a GetLogsBadRequest with default headers values
func NewGetLogsBadRequest() *GetLogsBadRequest {
	return &GetLogsBadRequest{}
}

/* GetLogsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetLogsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsBadRequest  %+v", 400, o.Payload)
}
func (o *GetLogsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/* GetLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLogsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLogsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogsInternalServerError creates a GetLogsInternalServerError with default headers values
func NewGetLogsInternalServerError() *GetLogsInternalServerError {
	return &GetLogsInternalServerError{}
}

/* GetLogsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLogsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLogsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
