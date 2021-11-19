// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/protodev-site/runtime"
	cr "github.com/protodev-site/runtime/client"
)

// NewFindTodosParams creates a new FindTodosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindTodosParams() *FindTodosParams {
	return &FindTodosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindTodosParamsWithTimeout creates a new FindTodosParams object
// with the ability to set a timeout on a request.
func NewFindTodosParamsWithTimeout(timeout time.Duration) *FindTodosParams {
	return &FindTodosParams{
		timeout: timeout,
	}
}

// NewFindTodosParamsWithContext creates a new FindTodosParams object
// with the ability to set a context for a request.
func NewFindTodosParamsWithContext(ctx context.Context) *FindTodosParams {
	return &FindTodosParams{
		Context: ctx,
	}
}

// NewFindTodosParamsWithHTTPClient creates a new FindTodosParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindTodosParamsWithHTTPClient(client *http.Client) *FindTodosParams {
	return &FindTodosParams{
		HTTPClient: client,
	}
}

/* FindTodosParams contains all the parameters to send to the API endpoint
   for the find todos operation.

   Typically these are written to a http.Request.
*/
type FindTodosParams struct {

	// Limit.
	//
	// Format: int32
	// Default: 20
	Limit *int32

	// Since.
	//
	// Format: int64
	Since *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find todos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindTodosParams) WithDefaults() *FindTodosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find todos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindTodosParams) SetDefaults() {
	var (
		limitDefault = int32(20)
	)

	val := FindTodosParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find todos params
func (o *FindTodosParams) WithTimeout(timeout time.Duration) *FindTodosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find todos params
func (o *FindTodosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find todos params
func (o *FindTodosParams) WithContext(ctx context.Context) *FindTodosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find todos params
func (o *FindTodosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find todos params
func (o *FindTodosParams) WithHTTPClient(client *http.Client) *FindTodosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find todos params
func (o *FindTodosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the find todos params
func (o *FindTodosParams) WithLimit(limit *int32) *FindTodosParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the find todos params
func (o *FindTodosParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the find todos params
func (o *FindTodosParams) WithSince(since *int64) *FindTodosParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the find todos params
func (o *FindTodosParams) SetSince(since *int64) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *FindTodosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Since != nil {

		// query param since
		var qrSince int64

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
