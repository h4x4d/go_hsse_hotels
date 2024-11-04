// Code generated by go-swagger; DO NOT EDIT.

package room

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/go_hsse_hotels/hotel/models"
)

// GetRoomByIDOKCode is the HTTP code returned for type GetRoomByIDOK
const GetRoomByIDOKCode int = 200

/*
GetRoomByIDOK successful operation

swagger:response getRoomByIdOK
*/
type GetRoomByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Room `json:"body,omitempty"`
}

// NewGetRoomByIDOK creates GetRoomByIDOK with default headers values
func NewGetRoomByIDOK() *GetRoomByIDOK {

	return &GetRoomByIDOK{}
}

// WithPayload adds the payload to the get room by Id o k response
func (o *GetRoomByIDOK) WithPayload(payload *models.Room) *GetRoomByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get room by Id o k response
func (o *GetRoomByIDOK) SetPayload(payload *models.Room) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoomByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRoomByIDNotFoundCode is the HTTP code returned for type GetRoomByIDNotFound
const GetRoomByIDNotFoundCode int = 404

/*
GetRoomByIDNotFound room not found

swagger:response getRoomByIdNotFound
*/
type GetRoomByIDNotFound struct {
}

// NewGetRoomByIDNotFound creates GetRoomByIDNotFound with default headers values
func NewGetRoomByIDNotFound() *GetRoomByIDNotFound {

	return &GetRoomByIDNotFound{}
}

// WriteResponse to the client
func (o *GetRoomByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
