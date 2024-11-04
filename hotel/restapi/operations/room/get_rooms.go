// Code generated by go-swagger; DO NOT EDIT.

package room

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRoomsHandlerFunc turns a function with the right signature into a get rooms handler
type GetRoomsHandlerFunc func(GetRoomsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoomsHandlerFunc) Handle(params GetRoomsParams) middleware.Responder {
	return fn(params)
}

// GetRoomsHandler interface for that can handle valid get rooms params
type GetRoomsHandler interface {
	Handle(GetRoomsParams) middleware.Responder
}

// NewGetRooms creates a new http.Handler for the get rooms operation
func NewGetRooms(ctx *middleware.Context, handler GetRoomsHandler) *GetRooms {
	return &GetRooms{Context: ctx, Handler: handler}
}

/*
	GetRooms swagger:route GET /hotel/room room getRooms

Get hotel rooms
*/
type GetRooms struct {
	Context *middleware.Context
	Handler GetRoomsHandler
}

func (o *GetRooms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRoomsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
