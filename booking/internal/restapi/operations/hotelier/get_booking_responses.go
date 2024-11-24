// Code generated by go-swagger; DO NOT EDIT.

package hotelier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

// GetBookingOKCode is the HTTP code returned for type GetBookingOK
const GetBookingOKCode int = 200

/*
GetBookingOK successful operation

swagger:response getBookingOK
*/
type GetBookingOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Booking `json:"body,omitempty"`
}

// NewGetBookingOK creates GetBookingOK with default headers values
func NewGetBookingOK() *GetBookingOK {

	return &GetBookingOK{}
}

// WithPayload adds the payload to the get booking o k response
func (o *GetBookingOK) WithPayload(payload []*models.Booking) *GetBookingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get booking o k response
func (o *GetBookingOK) SetPayload(payload []*models.Booking) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Booking, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBookingForbiddenCode is the HTTP code returned for type GetBookingForbidden
const GetBookingForbiddenCode int = 403

/*
GetBookingForbidden No access

swagger:response getBookingForbidden
*/
type GetBookingForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBookingForbidden creates GetBookingForbidden with default headers values
func NewGetBookingForbidden() *GetBookingForbidden {

	return &GetBookingForbidden{}
}

// WithPayload adds the payload to the get booking forbidden response
func (o *GetBookingForbidden) WithPayload(payload *models.Error) *GetBookingForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get booking forbidden response
func (o *GetBookingForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookingForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBookingNotFoundCode is the HTTP code returned for type GetBookingNotFound
const GetBookingNotFoundCode int = 404

/*
GetBookingNotFound Booking not found

swagger:response getBookingNotFound
*/
type GetBookingNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBookingNotFound creates GetBookingNotFound with default headers values
func NewGetBookingNotFound() *GetBookingNotFound {

	return &GetBookingNotFound{}
}

// WithPayload adds the payload to the get booking not found response
func (o *GetBookingNotFound) WithPayload(payload *models.Error) *GetBookingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get booking not found response
func (o *GetBookingNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}