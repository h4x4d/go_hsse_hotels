// Code generated by go-swagger; DO NOT EDIT.

package hotel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteHotelByIDOKCode is the HTTP code returned for type DeleteHotelByIDOK
const DeleteHotelByIDOKCode int = 200

/*
DeleteHotelByIDOK successful operation

swagger:response deleteHotelByIdOK
*/
type DeleteHotelByIDOK struct {
}

// NewDeleteHotelByIDOK creates DeleteHotelByIDOK with default headers values
func NewDeleteHotelByIDOK() *DeleteHotelByIDOK {

	return &DeleteHotelByIDOK{}
}

// WriteResponse to the client
func (o *DeleteHotelByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteHotelByIDNotFoundCode is the HTTP code returned for type DeleteHotelByIDNotFound
const DeleteHotelByIDNotFoundCode int = 404

/*
DeleteHotelByIDNotFound Hotel not found

swagger:response deleteHotelByIdNotFound
*/
type DeleteHotelByIDNotFound struct {
}

// NewDeleteHotelByIDNotFound creates DeleteHotelByIDNotFound with default headers values
func NewDeleteHotelByIDNotFound() *DeleteHotelByIDNotFound {

	return &DeleteHotelByIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteHotelByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
