// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/circl-dev/runtime"
)

// UploadFileOKCode is the HTTP code returned for type UploadFileOK
const UploadFileOKCode int = 200

/*UploadFileOK OK

swagger:response uploadFileOK
*/
type UploadFileOK struct {
}

// NewUploadFileOK creates UploadFileOK with default headers values
func NewUploadFileOK() *UploadFileOK {

	return &UploadFileOK{}
}

// WriteResponse to the client
func (o *UploadFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
