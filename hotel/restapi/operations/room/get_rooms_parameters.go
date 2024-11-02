// Code generated by go-swagger; DO NOT EDIT.

package room

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRoomsParams creates a new GetRoomsParams object
//
// There are no default values defined in the spec.
func NewGetRoomsParams() GetRoomsParams {

	return GetRoomsParams{}
}

// GetRoomsParams contains all the bound params for the get rooms operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_rooms
type GetRoomsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	HotelID *int64
	/*
	  In: query
	*/
	Tag *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRoomsParams() beforehand.
func (o *GetRoomsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHotelID, qhkHotelID, _ := qs.GetOK("hotel_id")
	if err := o.bindHotelID(qHotelID, qhkHotelID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTag, qhkTag, _ := qs.GetOK("tag")
	if err := o.bindTag(qTag, qhkTag, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHotelID binds and validates parameter HotelID from query.
func (o *GetRoomsParams) bindHotelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("hotel_id", "query", "int64", raw)
	}
	o.HotelID = &value

	return nil
}

// bindTag binds and validates parameter Tag from query.
func (o *GetRoomsParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Tag = &raw

	return nil
}