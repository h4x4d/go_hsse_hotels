// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"auth/internal/models"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*
PostLoginOK Success

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload *PostLoginOKBody `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload *PostLoginOKBody) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload *PostLoginOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginUnauthorizedCode is the HTTP code returned for type PostLoginUnauthorized
const PostLoginUnauthorizedCode int = 401

/*
PostLoginUnauthorized Incorrect login data

swagger:response postLoginUnauthorized
*/
type PostLoginUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostLoginUnauthorized creates PostLoginUnauthorized with default headers values
func NewPostLoginUnauthorized() *PostLoginUnauthorized {

	return &PostLoginUnauthorized{}
}

// WithPayload adds the payload to the post login unauthorized response
func (o *PostLoginUnauthorized) WithPayload(payload *models.Error) *PostLoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login unauthorized response
func (o *PostLoginUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
