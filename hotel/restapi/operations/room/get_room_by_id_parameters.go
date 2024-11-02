// Code generated by go-swagger; DO NOT EDIT.

package room

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRoomByIDParams creates a new GetRoomByIDParams object
//
// There are no default values defined in the spec.
func NewGetRoomByIDParams() GetRoomByIDParams {

	return GetRoomByIDParams{}
}

// GetRoomByIDParams contains all the bound params for the get room by id operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_room_by_id
type GetRoomByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of room to return
	  Required: true
	  In: path
	*/
	RoomID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRoomByIDParams() beforehand.
func (o *GetRoomByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRoomID, rhkRoomID, _ := route.Params.GetOK("room_id")
	if err := o.bindRoomID(rRoomID, rhkRoomID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRoomID binds and validates parameter RoomID from path.
func (o *GetRoomByIDParams) bindRoomID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("room_id", "path", "int64", raw)
	}
	o.RoomID = value

	return nil
}
