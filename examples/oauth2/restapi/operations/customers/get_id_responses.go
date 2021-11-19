// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/protodev-site/runtime"

	"github.com/protodev-site/go-swagger/examples/oauth2/models"
)

// GetIDOKCode is the HTTP code returned for type GetIDOK
const GetIDOKCode int = 200

/*GetIDOK OK

swagger:response getIdOK
*/
type GetIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Customer `json:"body,omitempty"`
}

// NewGetIDOK creates GetIDOK with default headers values
func NewGetIDOK() *GetIDOK {

	return &GetIDOK{}
}

// WithPayload adds the payload to the get Id o k response
func (o *GetIDOK) WithPayload(payload *models.Customer) *GetIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id o k response
func (o *GetIDOK) SetPayload(payload *models.Customer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIDUnauthorizedCode is the HTTP code returned for type GetIDUnauthorized
const GetIDUnauthorizedCode int = 401

/*GetIDUnauthorized unauthorized

swagger:response getIdUnauthorized
*/
type GetIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIDUnauthorized creates GetIDUnauthorized with default headers values
func NewGetIDUnauthorized() *GetIDUnauthorized {

	return &GetIDUnauthorized{}
}

// WithPayload adds the payload to the get Id unauthorized response
func (o *GetIDUnauthorized) WithPayload(payload *models.Error) *GetIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id unauthorized response
func (o *GetIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIDNotFoundCode is the HTTP code returned for type GetIDNotFound
const GetIDNotFoundCode int = 404

/*GetIDNotFound resource not found

swagger:response getIdNotFound
*/
type GetIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIDNotFound creates GetIDNotFound with default headers values
func NewGetIDNotFound() *GetIDNotFound {

	return &GetIDNotFound{}
}

// WithPayload adds the payload to the get Id not found response
func (o *GetIDNotFound) WithPayload(payload *models.Error) *GetIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id not found response
func (o *GetIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetIDDefault error

swagger:response getIdDefault
*/
type GetIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIDDefault creates GetIDDefault with default headers values
func NewGetIDDefault(code int) *GetIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get Id default response
func (o *GetIDDefault) WithStatusCode(code int) *GetIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get Id default response
func (o *GetIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get Id default response
func (o *GetIDDefault) WithPayload(payload *models.Error) *GetIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id default response
func (o *GetIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
