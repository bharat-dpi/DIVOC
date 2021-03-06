// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/api/swagger_gen/models"
)

// UpdateCertificateOKCode is the HTTP code returned for type UpdateCertificateOK
const UpdateCertificateOKCode int = 200

/*UpdateCertificateOK OK

swagger:response updateCertificateOK
*/
type UpdateCertificateOK struct {
}

// NewUpdateCertificateOK creates UpdateCertificateOK with default headers values
func NewUpdateCertificateOK() *UpdateCertificateOK {

	return &UpdateCertificateOK{}
}

// WriteResponse to the client
func (o *UpdateCertificateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateCertificateBadRequestCode is the HTTP code returned for type UpdateCertificateBadRequest
const UpdateCertificateBadRequestCode int = 400

/*UpdateCertificateBadRequest Invalid input

swagger:response updateCertificateBadRequest
*/
type UpdateCertificateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCertificateBadRequest creates UpdateCertificateBadRequest with default headers values
func NewUpdateCertificateBadRequest() *UpdateCertificateBadRequest {

	return &UpdateCertificateBadRequest{}
}

// WithPayload adds the payload to the update certificate bad request response
func (o *UpdateCertificateBadRequest) WithPayload(payload *models.Error) *UpdateCertificateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update certificate bad request response
func (o *UpdateCertificateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCertificateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCertificatePreconditionFailedCode is the HTTP code returned for type UpdateCertificatePreconditionFailed
const UpdateCertificatePreconditionFailedCode int = 412

/*UpdateCertificatePreconditionFailed Update not allowed

swagger:response updateCertificatePreconditionFailed
*/
type UpdateCertificatePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCertificatePreconditionFailed creates UpdateCertificatePreconditionFailed with default headers values
func NewUpdateCertificatePreconditionFailed() *UpdateCertificatePreconditionFailed {

	return &UpdateCertificatePreconditionFailed{}
}

// WithPayload adds the payload to the update certificate precondition failed response
func (o *UpdateCertificatePreconditionFailed) WithPayload(payload *models.Error) *UpdateCertificatePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update certificate precondition failed response
func (o *UpdateCertificatePreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCertificatePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
