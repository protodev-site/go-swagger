// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/circl-dev/runtime"

	"github.com/circl-dev/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetGetOKCode is the HTTP code returned for type PetGetOK
const PetGetOKCode int = 200

/*PetGetOK successful operation

swagger:response petGetOK
*/
type PetGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewPetGetOK creates PetGetOK with default headers values
func NewPetGetOK() *PetGetOK {

	return &PetGetOK{}
}

// WithPayload adds the payload to the pet get o k response
func (o *PetGetOK) WithPayload(payload *models.Pet) *PetGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pet get o k response
func (o *PetGetOK) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PetGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PetGetBadRequestCode is the HTTP code returned for type PetGetBadRequest
const PetGetBadRequestCode int = 400

/*PetGetBadRequest Invalid ID supplied

swagger:response petGetBadRequest
*/
type PetGetBadRequest struct {
}

// NewPetGetBadRequest creates PetGetBadRequest with default headers values
func NewPetGetBadRequest() *PetGetBadRequest {

	return &PetGetBadRequest{}
}

// WriteResponse to the client
func (o *PetGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PetGetNotFoundCode is the HTTP code returned for type PetGetNotFound
const PetGetNotFoundCode int = 404

/*PetGetNotFound Pet not found

swagger:response petGetNotFound
*/
type PetGetNotFound struct {
}

// NewPetGetNotFound creates PetGetNotFound with default headers values
func NewPetGetNotFound() *PetGetNotFound {

	return &PetGetNotFound{}
}

// WriteResponse to the client
func (o *PetGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
