// Code generated by go-swagger; DO NOT EDIT.

package certificate_revoked

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CertificateRevokedOKCode is the HTTP code returned for type CertificateRevokedOK
const CertificateRevokedOKCode int = 200

/*CertificateRevokedOK OK

swagger:response certificateRevokedOK
*/
type CertificateRevokedOK struct {
}

// NewCertificateRevokedOK creates CertificateRevokedOK with default headers values
func NewCertificateRevokedOK() *CertificateRevokedOK {

	return &CertificateRevokedOK{}
}

// WriteResponse to the client
func (o *CertificateRevokedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CertificateRevokedBadRequestCode is the HTTP code returned for type CertificateRevokedBadRequest
const CertificateRevokedBadRequestCode int = 400

/*CertificateRevokedBadRequest Invalid input

swagger:response certificateRevokedBadRequest
*/
type CertificateRevokedBadRequest struct {
}

// NewCertificateRevokedBadRequest creates CertificateRevokedBadRequest with default headers values
func NewCertificateRevokedBadRequest() *CertificateRevokedBadRequest {

	return &CertificateRevokedBadRequest{}
}

// WriteResponse to the client
func (o *CertificateRevokedBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CertificateRevokedNotFoundCode is the HTTP code returned for type CertificateRevokedNotFound
const CertificateRevokedNotFoundCode int = 404

/*CertificateRevokedNotFound certificate not found in revocation list

swagger:response certificateRevokedNotFound
*/
type CertificateRevokedNotFound struct {
}

// NewCertificateRevokedNotFound creates CertificateRevokedNotFound with default headers values
func NewCertificateRevokedNotFound() *CertificateRevokedNotFound {

	return &CertificateRevokedNotFound{}
}

// WriteResponse to the client
func (o *CertificateRevokedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
