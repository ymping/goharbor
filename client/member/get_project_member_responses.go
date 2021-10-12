// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ymping/goharbor/models"
)

// GetProjectMemberReader is a Reader for the GetProjectMember structure.
type GetProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectMemberOK creates a GetProjectMemberOK with default headers values
func NewGetProjectMemberOK() *GetProjectMemberOK {
	return &GetProjectMemberOK{}
}

/* GetProjectMemberOK describes a response with status code 200, with default header values.

Project member retrieved successfully.
*/
type GetProjectMemberOK struct {
	Payload *models.ProjectMemberEntity
}

func (o *GetProjectMemberOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberOK  %+v", 200, o.Payload)
}
func (o *GetProjectMemberOK) GetPayload() *models.ProjectMemberEntity {
	return o.Payload
}

func (o *GetProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectMemberEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMemberBadRequest creates a GetProjectMemberBadRequest with default headers values
func NewGetProjectMemberBadRequest() *GetProjectMemberBadRequest {
	return &GetProjectMemberBadRequest{}
}

/* GetProjectMemberBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetProjectMemberBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberBadRequest  %+v", 400, o.Payload)
}
func (o *GetProjectMemberBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberUnauthorized creates a GetProjectMemberUnauthorized with default headers values
func NewGetProjectMemberUnauthorized() *GetProjectMemberUnauthorized {
	return &GetProjectMemberUnauthorized{}
}

/* GetProjectMemberUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectMemberUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProjectMemberUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberForbidden creates a GetProjectMemberForbidden with default headers values
func NewGetProjectMemberForbidden() *GetProjectMemberForbidden {
	return &GetProjectMemberForbidden{}
}

/* GetProjectMemberForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectMemberForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberForbidden  %+v", 403, o.Payload)
}
func (o *GetProjectMemberForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberNotFound creates a GetProjectMemberNotFound with default headers values
func NewGetProjectMemberNotFound() *GetProjectMemberNotFound {
	return &GetProjectMemberNotFound{}
}

/* GetProjectMemberNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProjectMemberNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberNotFound  %+v", 404, o.Payload)
}
func (o *GetProjectMemberNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberInternalServerError creates a GetProjectMemberInternalServerError with default headers values
func NewGetProjectMemberInternalServerError() *GetProjectMemberInternalServerError {
	return &GetProjectMemberInternalServerError{}
}

/* GetProjectMemberInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetProjectMemberInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProjectMemberInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
