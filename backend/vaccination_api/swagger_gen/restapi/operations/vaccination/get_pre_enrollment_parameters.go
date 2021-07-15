// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetPreEnrollmentParams creates a new GetPreEnrollmentParams object
//
// There are no default values defined in the spec.
func NewGetPreEnrollmentParams() GetPreEnrollmentParams {

	return GetPreEnrollmentParams{}
}

// GetPreEnrollmentParams contains all the bound params for the get pre enrollment operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPreEnrollment
type GetPreEnrollmentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	PreEnrollmentCode string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPreEnrollmentParams() beforehand.
func (o *GetPreEnrollmentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPreEnrollmentCode, rhkPreEnrollmentCode, _ := route.Params.GetOK("preEnrollmentCode")
	if err := o.bindPreEnrollmentCode(rPreEnrollmentCode, rhkPreEnrollmentCode, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPreEnrollmentCode binds and validates parameter PreEnrollmentCode from path.
func (o *GetPreEnrollmentParams) bindPreEnrollmentCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PreEnrollmentCode = raw

	return nil
}
