// Code generated by go-swagger; DO NOT EDIT.

package room

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
)

// NewUpdateRoomParams creates a new UpdateRoomParams object
//
// There are no default values defined in the spec.
func NewUpdateRoomParams() UpdateRoomParams {

	return UpdateRoomParams{}
}

// UpdateRoomParams contains all the bound params for the update room operation
// typically these are obtained from a http.Request
//
// swagger:parameters update_room
type UpdateRoomParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Object *models.Room
	/*ID of room to change
	  Required: true
	  In: path
	*/
	RoomID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateRoomParams() beforehand.
func (o *UpdateRoomParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Room
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("object", "body", ""))
			} else {
				res = append(res, errors.NewParseError("object", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Object = &body
			}
		}
	} else {
		res = append(res, errors.Required("object", "body", ""))
	}

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
func (o *UpdateRoomParams) bindRoomID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
