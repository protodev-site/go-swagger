// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/circl-dev/runtime"
	cr "github.com/circl-dev/runtime/client"
)

// NewPetListParams creates a new PetListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPetListParams() *PetListParams {
	return &PetListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPetListParamsWithTimeout creates a new PetListParams object
// with the ability to set a timeout on a request.
func NewPetListParamsWithTimeout(timeout time.Duration) *PetListParams {
	return &PetListParams{
		timeout: timeout,
	}
}

// NewPetListParamsWithContext creates a new PetListParams object
// with the ability to set a context for a request.
func NewPetListParamsWithContext(ctx context.Context) *PetListParams {
	return &PetListParams{
		Context: ctx,
	}
}

// NewPetListParamsWithHTTPClient creates a new PetListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPetListParamsWithHTTPClient(client *http.Client) *PetListParams {
	return &PetListParams{
		HTTPClient: client,
	}
}

/* PetListParams contains all the parameters to send to the API endpoint
   for the pet list operation.

   Typically these are written to a http.Request.
*/
type PetListParams struct {

	/* Status.

	   Status values that need to be considered for filter
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pet list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PetListParams) WithDefaults() *PetListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pet list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PetListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pet list params
func (o *PetListParams) WithTimeout(timeout time.Duration) *PetListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pet list params
func (o *PetListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pet list params
func (o *PetListParams) WithContext(ctx context.Context) *PetListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pet list params
func (o *PetListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pet list params
func (o *PetListParams) WithHTTPClient(client *http.Client) *PetListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pet list params
func (o *PetListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatus adds the status to the pet list params
func (o *PetListParams) WithStatus(status []string) *PetListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the pet list params
func (o *PetListParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *PetListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPetList binds the parameter status
func (o *PetListParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "multi"
	statusIS := swag.JoinByFormat(statusIC, "multi")

	return statusIS
}
