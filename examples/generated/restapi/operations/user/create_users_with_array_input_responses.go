// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/circl-dev/runtime"
)

/*CreateUsersWithArrayInputDefault successful operation

swagger:response createUsersWithArrayInputDefault
*/
type CreateUsersWithArrayInputDefault struct {
	_statusCode int
}

// NewCreateUsersWithArrayInputDefault creates CreateUsersWithArrayInputDefault with default headers values
func NewCreateUsersWithArrayInputDefault(code int) *CreateUsersWithArrayInputDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUsersWithArrayInputDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create users with array input default response
func (o *CreateUsersWithArrayInputDefault) WithStatusCode(code int) *CreateUsersWithArrayInputDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create users with array input default response
func (o *CreateUsersWithArrayInputDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *CreateUsersWithArrayInputDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
