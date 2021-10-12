// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ymping/goharbor/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/* DeleteUserOK describes a response with status code 200, with default header values.

Success
*/
type DeleteUserOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserOK ", 200)
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteUserUnauthorized creates a DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

/* DeleteUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUserUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteUserUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserForbidden creates a DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {
	return &DeleteUserForbidden{}
}

/* DeleteUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUserForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserForbidden  %+v", 403, o.Payload)
}
func (o *DeleteUserForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/* DeleteUserNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteUserNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserInternalServerError creates a DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

/* DeleteUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteUserInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteUserInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
