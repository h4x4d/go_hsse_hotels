// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateBookingHandlerFunc turns a function with the right signature into a create booking handler
type CreateBookingHandlerFunc func(CreateBookingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBookingHandlerFunc) Handle(params CreateBookingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateBookingHandler interface for that can handle valid create booking params
type CreateBookingHandler interface {
	Handle(CreateBookingParams, interface{}) middleware.Responder
}

// NewCreateBooking creates a new http.Handler for the create booking operation
func NewCreateBooking(ctx *middleware.Context, handler CreateBookingHandler) *CreateBooking {
	return &CreateBooking{Context: ctx, Handler: handler}
}

/*
	CreateBooking swagger:route POST /booking customer createBooking

Create booking
*/
type CreateBooking struct {
	Context *middleware.Context
	Handler CreateBookingHandler
}

func (o *CreateBooking) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateBookingParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateBookingOKBody create booking o k body
//
// swagger:model CreateBookingOKBody
type CreateBookingOKBody struct {

	// booking id
	BookingID int64 `json:"booking_id,omitempty"`
}

// Validate validates this create booking o k body
func (o *CreateBookingOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create booking o k body based on context it is used
func (o *CreateBookingOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBookingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBookingOKBody) UnmarshalBinary(b []byte) error {
	var res CreateBookingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}