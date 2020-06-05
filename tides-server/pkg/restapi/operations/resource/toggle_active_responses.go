// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ToggleActiveOKCode is the HTTP code returned for type ToggleActiveOK
const ToggleActiveOKCode int = 200

/*ToggleActiveOK returns success message

swagger:response toggleActiveOK
*/
type ToggleActiveOK struct {

	/*
	  In: Body
	*/
	Payload *ToggleActiveOKBody `json:"body,omitempty"`
}

// NewToggleActiveOK creates ToggleActiveOK with default headers values
func NewToggleActiveOK() *ToggleActiveOK {

	return &ToggleActiveOK{}
}

// WithPayload adds the payload to the toggle active o k response
func (o *ToggleActiveOK) WithPayload(payload *ToggleActiveOKBody) *ToggleActiveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the toggle active o k response
func (o *ToggleActiveOK) SetPayload(payload *ToggleActiveOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ToggleActiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ToggleActiveUnauthorizedCode is the HTTP code returned for type ToggleActiveUnauthorized
const ToggleActiveUnauthorizedCode int = 401

/*ToggleActiveUnauthorized Unauthorized

swagger:response toggleActiveUnauthorized
*/
type ToggleActiveUnauthorized struct {
}

// NewToggleActiveUnauthorized creates ToggleActiveUnauthorized with default headers values
func NewToggleActiveUnauthorized() *ToggleActiveUnauthorized {

	return &ToggleActiveUnauthorized{}
}

// WriteResponse to the client
func (o *ToggleActiveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ToggleActiveNotFoundCode is the HTTP code returned for type ToggleActiveNotFound
const ToggleActiveNotFoundCode int = 404

/*ToggleActiveNotFound resouce not found

swagger:response toggleActiveNotFound
*/
type ToggleActiveNotFound struct {

	/*
	  In: Body
	*/
	Payload *ToggleActiveNotFoundBody `json:"body,omitempty"`
}

// NewToggleActiveNotFound creates ToggleActiveNotFound with default headers values
func NewToggleActiveNotFound() *ToggleActiveNotFound {

	return &ToggleActiveNotFound{}
}

// WithPayload adds the payload to the toggle active not found response
func (o *ToggleActiveNotFound) WithPayload(payload *ToggleActiveNotFoundBody) *ToggleActiveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the toggle active not found response
func (o *ToggleActiveNotFound) SetPayload(payload *ToggleActiveNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ToggleActiveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
