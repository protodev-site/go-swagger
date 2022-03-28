// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/circl-dev/runtime"

	custom "github.com/circl-dev/go-swagger/examples/external-types/fred"
)

/*PostTestDefault An inlined reference to an aliased external package.
The response is defined as map[string][]map[string]custom.MyAlternateString

No definition is generated in models.


swagger:response postTestDefault
*/
type PostTestDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload map[string][]map[string]custom.MyAlternateString `json:"body,omitempty"`
}

// NewPostTestDefault creates PostTestDefault with default headers values
func NewPostTestDefault(code int) *PostTestDefault {
	if code <= 0 {
		code = 500
	}

	return &PostTestDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post test default response
func (o *PostTestDefault) WithStatusCode(code int) *PostTestDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post test default response
func (o *PostTestDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post test default response
func (o *PostTestDefault) WithPayload(payload map[string][]map[string]custom.MyAlternateString) *PostTestDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test default response
func (o *PostTestDefault) SetPayload(payload map[string][]map[string]custom.MyAlternateString) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string][]map[string]custom.MyAlternateString, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
