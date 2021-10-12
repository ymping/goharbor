// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ymping/goharbor/models"
)

// GetRobotByIDReader is a Reader for the GetRobotByID structure.
type GetRobotByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRobotByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRobotByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRobotByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRobotByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRobotByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRobotByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRobotByIDOK creates a GetRobotByIDOK with default headers values
func NewGetRobotByIDOK() *GetRobotByIDOK {
	return &GetRobotByIDOK{}
}

/* GetRobotByIDOK describes a response with status code 200, with default header values.

Return matched robot information.
*/
type GetRobotByIDOK struct {
	Payload *models.Robot
}

func (o *GetRobotByIDOK) Error() string {
	return fmt.Sprintf("[GET /robots/{robot_id}][%d] getRobotByIdOK  %+v", 200, o.Payload)
}
func (o *GetRobotByIDOK) GetPayload() *models.Robot {
	return o.Payload
}

func (o *GetRobotByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Robot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRobotByIDUnauthorized creates a GetRobotByIDUnauthorized with default headers values
func NewGetRobotByIDUnauthorized() *GetRobotByIDUnauthorized {
	return &GetRobotByIDUnauthorized{}
}

/* GetRobotByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRobotByIDUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRobotByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /robots/{robot_id}][%d] getRobotByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRobotByIDUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDForbidden creates a GetRobotByIDForbidden with default headers values
func NewGetRobotByIDForbidden() *GetRobotByIDForbidden {
	return &GetRobotByIDForbidden{}
}

/* GetRobotByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRobotByIDForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRobotByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /robots/{robot_id}][%d] getRobotByIdForbidden  %+v", 403, o.Payload)
}
func (o *GetRobotByIDForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDNotFound creates a GetRobotByIDNotFound with default headers values
func NewGetRobotByIDNotFound() *GetRobotByIDNotFound {
	return &GetRobotByIDNotFound{}
}

/* GetRobotByIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRobotByIDNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRobotByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /robots/{robot_id}][%d] getRobotByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetRobotByIDNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRobotByIDInternalServerError creates a GetRobotByIDInternalServerError with default headers values
func NewGetRobotByIDInternalServerError() *GetRobotByIDInternalServerError {
	return &GetRobotByIDInternalServerError{}
}

/* GetRobotByIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRobotByIDInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetRobotByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /robots/{robot_id}][%d] getRobotByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRobotByIDInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRobotByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
