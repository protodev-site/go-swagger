// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime"

	"github.com/circl-dev/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetUpdateReader is a Reader for the PetUpdate structure.
type PetUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPetUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPetUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPetUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPetUpdateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPetUpdateCreated creates a PetUpdateCreated with default headers values
func NewPetUpdateCreated() *PetUpdateCreated {
	return &PetUpdateCreated{}
}

/* PetUpdateCreated describes a response with status code 201, with default header values.

Updated successfully
*/
type PetUpdateCreated struct {
	Payload *models.Pet
}

func (o *PetUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateCreated  %+v", 201, o.Payload)
}
func (o *PetUpdateCreated) GetPayload() *models.Pet {
	return o.Payload
}

func (o *PetUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetUpdateBadRequest creates a PetUpdateBadRequest with default headers values
func NewPetUpdateBadRequest() *PetUpdateBadRequest {
	return &PetUpdateBadRequest{}
}

/* PetUpdateBadRequest describes a response with status code 400, with default header values.

Invalid ID supplied
*/
type PetUpdateBadRequest struct {
}

func (o *PetUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateBadRequest ", 400)
}

func (o *PetUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetUpdateNotFound creates a PetUpdateNotFound with default headers values
func NewPetUpdateNotFound() *PetUpdateNotFound {
	return &PetUpdateNotFound{}
}

/* PetUpdateNotFound describes a response with status code 404, with default header values.

Pet not found
*/
type PetUpdateNotFound struct {
}

func (o *PetUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateNotFound ", 404)
}

func (o *PetUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetUpdateMethodNotAllowed creates a PetUpdateMethodNotAllowed with default headers values
func NewPetUpdateMethodNotAllowed() *PetUpdateMethodNotAllowed {
	return &PetUpdateMethodNotAllowed{}
}

/* PetUpdateMethodNotAllowed describes a response with status code 405, with default header values.

Validation exception
*/
type PetUpdateMethodNotAllowed struct {
}

func (o *PetUpdateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] petUpdateMethodNotAllowed ", 405)
}

func (o *PetUpdateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
