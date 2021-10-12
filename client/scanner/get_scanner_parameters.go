// Code generated by go-swagger; DO NOT EDIT.

package scanner

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

// NewGetScannerParams creates a new GetScannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScannerParams() *GetScannerParams {
	return &GetScannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScannerParamsWithTimeout creates a new GetScannerParams object
// with the ability to set a timeout on a request.
func NewGetScannerParamsWithTimeout(timeout time.Duration) *GetScannerParams {
	return &GetScannerParams{
		timeout: timeout,
	}
}

// NewGetScannerParamsWithContext creates a new GetScannerParams object
// with the ability to set a context for a request.
func NewGetScannerParamsWithContext(ctx context.Context) *GetScannerParams {
	return &GetScannerParams{
		Context: ctx,
	}
}

// NewGetScannerParamsWithHTTPClient creates a new GetScannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScannerParamsWithHTTPClient(client *http.Client) *GetScannerParams {
	return &GetScannerParams{
		HTTPClient: client,
	}
}

/* GetScannerParams contains all the parameters to send to the API endpoint
   for the get scanner operation.

   Typically these are written to a http.Request.
*/
type GetScannerParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* RegistrationID.

	   The scanner registration identifer.
	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScannerParams) WithDefaults() *GetScannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scanner params
func (o *GetScannerParams) WithTimeout(timeout time.Duration) *GetScannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scanner params
func (o *GetScannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scanner params
func (o *GetScannerParams) WithContext(ctx context.Context) *GetScannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scanner params
func (o *GetScannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scanner params
func (o *GetScannerParams) WithHTTPClient(client *http.Client) *GetScannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scanner params
func (o *GetScannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get scanner params
func (o *GetScannerParams) WithXRequestID(xRequestID *string) *GetScannerParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get scanner params
func (o *GetScannerParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistrationID adds the registrationID to the get scanner params
func (o *GetScannerParams) WithRegistrationID(registrationID string) *GetScannerParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the get scanner params
func (o *GetScannerParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
