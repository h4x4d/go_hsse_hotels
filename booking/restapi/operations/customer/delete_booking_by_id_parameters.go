// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteBookingByIDParams creates a new DeleteBookingByIDParams object
//
// There are no default values defined in the spec.
func NewDeleteBookingByIDParams() DeleteBookingByIDParams {

	return DeleteBookingByIDParams{}
}

// DeleteBookingByIDParams contains all the bound params for the delete booking by id operation
// typically these are obtained from a http.Request
//
// swagger:parameters delete_booking_by_id
type DeleteBookingByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of booking to delete
	  Required: true
	  In: path
	*/
	BookingID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteBookingByIDParams() beforehand.
func (o *DeleteBookingByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBookingID, rhkBookingID, _ := route.Params.GetOK("booking_id")
	if err := o.bindBookingID(rBookingID, rhkBookingID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBookingID binds and validates parameter BookingID from path.
func (o *DeleteBookingByIDParams) bindBookingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("booking_id", "path", "int64", raw)
	}
	o.BookingID = value

	return nil
}
