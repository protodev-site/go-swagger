// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/circl-dev/runtime"

	"github.com/circl-dev/go-swagger/examples/composed-auth/models"
)

// GetAccountOKCode is the HTTP code returned for type GetAccountOK
const GetAccountOKCode int = 200

/*GetAccountOK registered user personal account infos

swagger:response getAccountOK
*/
type GetAccountOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAccountOK creates GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {

	return &GetAccountOK{}
}

// WithPayload adds the payload to the get account o k response
func (o *GetAccountOK) WithPayload(payload interface{}) *GetAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account o k response
func (o *GetAccountOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountUnauthorizedCode is the HTTP code returned for type GetAccountUnauthorized
const GetAccountUnauthorizedCode int = 401

/*GetAccountUnauthorized unauthorized access for a lack of authentication

swagger:response getAccountUnauthorized
*/
type GetAccountUnauthorized struct {
}

// NewGetAccountUnauthorized creates GetAccountUnauthorized with default headers values
func NewGetAccountUnauthorized() *GetAccountUnauthorized {

	return &GetAccountUnauthorized{}
}

// WriteResponse to the client
func (o *GetAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

/*GetAccountDefault other error response

swagger:response getAccountDefault
*/
type GetAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAccountDefault creates GetAccountDefault with default headers values
func NewGetAccountDefault(code int) *GetAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get account default response
func (o *GetAccountDefault) WithStatusCode(code int) *GetAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get account default response
func (o *GetAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get account default response
func (o *GetAccountDefault) WithPayload(payload *models.Error) *GetAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account default response
func (o *GetAccountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
