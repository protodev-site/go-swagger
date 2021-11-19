// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/protodev-site/runtime"

	"github.com/protodev-site/go-swagger/examples/composed-auth/models"
)

// AddOrderOKCode is the HTTP code returned for type AddOrderOK
const AddOrderOKCode int = 200

/*AddOrderOK empty response

swagger:response addOrderOK
*/
type AddOrderOK struct {
}

// NewAddOrderOK creates AddOrderOK with default headers values
func NewAddOrderOK() *AddOrderOK {

	return &AddOrderOK{}
}

// WriteResponse to the client
func (o *AddOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddOrderUnauthorizedCode is the HTTP code returned for type AddOrderUnauthorized
const AddOrderUnauthorizedCode int = 401

/*AddOrderUnauthorized unauthorized access for a lack of authentication

swagger:response addOrderUnauthorized
*/
type AddOrderUnauthorized struct {
}

// NewAddOrderUnauthorized creates AddOrderUnauthorized with default headers values
func NewAddOrderUnauthorized() *AddOrderUnauthorized {

	return &AddOrderUnauthorized{}
}

// WriteResponse to the client
func (o *AddOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AddOrderForbiddenCode is the HTTP code returned for type AddOrderForbidden
const AddOrderForbiddenCode int = 403

/*AddOrderForbidden forbidden access for a lack of sufficient privileges

swagger:response addOrderForbidden
*/
type AddOrderForbidden struct {
}

// NewAddOrderForbidden creates AddOrderForbidden with default headers values
func NewAddOrderForbidden() *AddOrderForbidden {

	return &AddOrderForbidden{}
}

// WriteResponse to the client
func (o *AddOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*AddOrderDefault other error response

swagger:response addOrderDefault
*/
type AddOrderDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddOrderDefault creates AddOrderDefault with default headers values
func NewAddOrderDefault(code int) *AddOrderDefault {
	if code <= 0 {
		code = 500
	}

	return &AddOrderDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add order default response
func (o *AddOrderDefault) WithStatusCode(code int) *AddOrderDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add order default response
func (o *AddOrderDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add order default response
func (o *AddOrderDefault) WithPayload(payload *models.Error) *AddOrderDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add order default response
func (o *AddOrderDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOrderDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
