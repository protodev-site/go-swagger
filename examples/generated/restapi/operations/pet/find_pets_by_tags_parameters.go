// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime"
	"github.com/circl-dev/runtime/middleware"
)

// NewFindPetsByTagsParams creates a new FindPetsByTagsParams object
//
// There are no default values defined in the spec.
func NewFindPetsByTagsParams() FindPetsByTagsParams {

	return FindPetsByTagsParams{}
}

// FindPetsByTagsParams contains all the bound params for the find pets by tags operation
// typically these are obtained from a http.Request
//
// swagger:parameters findPetsByTags
type FindPetsByTagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Tags to filter by
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindPetsByTagsParams() beforehand.
func (o *FindPetsByTagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *FindPetsByTagsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	// CollectionFormat: multi
	tagsIC := rawData
	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
