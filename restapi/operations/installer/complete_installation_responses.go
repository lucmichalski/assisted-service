// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/filanov/bm-inventory/models"
)

// CompleteInstallationAcceptedCode is the HTTP code returned for type CompleteInstallationAccepted
const CompleteInstallationAcceptedCode int = 202

/*CompleteInstallationAccepted Success.

swagger:response completeInstallationAccepted
*/
type CompleteInstallationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewCompleteInstallationAccepted creates CompleteInstallationAccepted with default headers values
func NewCompleteInstallationAccepted() *CompleteInstallationAccepted {

	return &CompleteInstallationAccepted{}
}

// WithPayload adds the payload to the complete installation accepted response
func (o *CompleteInstallationAccepted) WithPayload(payload *models.Cluster) *CompleteInstallationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the complete installation accepted response
func (o *CompleteInstallationAccepted) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompleteInstallationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompleteInstallationNotFoundCode is the HTTP code returned for type CompleteInstallationNotFound
const CompleteInstallationNotFoundCode int = 404

/*CompleteInstallationNotFound Error.

swagger:response completeInstallationNotFound
*/
type CompleteInstallationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCompleteInstallationNotFound creates CompleteInstallationNotFound with default headers values
func NewCompleteInstallationNotFound() *CompleteInstallationNotFound {

	return &CompleteInstallationNotFound{}
}

// WithPayload adds the payload to the complete installation not found response
func (o *CompleteInstallationNotFound) WithPayload(payload *models.Error) *CompleteInstallationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the complete installation not found response
func (o *CompleteInstallationNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompleteInstallationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompleteInstallationConflictCode is the HTTP code returned for type CompleteInstallationConflict
const CompleteInstallationConflictCode int = 409

/*CompleteInstallationConflict Error.

swagger:response completeInstallationConflict
*/
type CompleteInstallationConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCompleteInstallationConflict creates CompleteInstallationConflict with default headers values
func NewCompleteInstallationConflict() *CompleteInstallationConflict {

	return &CompleteInstallationConflict{}
}

// WithPayload adds the payload to the complete installation conflict response
func (o *CompleteInstallationConflict) WithPayload(payload *models.Error) *CompleteInstallationConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the complete installation conflict response
func (o *CompleteInstallationConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompleteInstallationConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompleteInstallationInternalServerErrorCode is the HTTP code returned for type CompleteInstallationInternalServerError
const CompleteInstallationInternalServerErrorCode int = 500

/*CompleteInstallationInternalServerError Error.

swagger:response completeInstallationInternalServerError
*/
type CompleteInstallationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCompleteInstallationInternalServerError creates CompleteInstallationInternalServerError with default headers values
func NewCompleteInstallationInternalServerError() *CompleteInstallationInternalServerError {

	return &CompleteInstallationInternalServerError{}
}

// WithPayload adds the payload to the complete installation internal server error response
func (o *CompleteInstallationInternalServerError) WithPayload(payload *models.Error) *CompleteInstallationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the complete installation internal server error response
func (o *CompleteInstallationInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompleteInstallationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}