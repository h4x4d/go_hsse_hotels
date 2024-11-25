// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostRegisterHandlerFunc turns a function with the right signature into a post register handler
type PostRegisterHandlerFunc func(PostRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRegisterHandlerFunc) Handle(params PostRegisterParams) middleware.Responder {
	return fn(params)
}

// PostRegisterHandler interface for that can handle valid post register params
type PostRegisterHandler interface {
	Handle(PostRegisterParams) middleware.Responder
}

// NewPostRegister creates a new http.Handler for the post register operation
func NewPostRegister(ctx *middleware.Context, handler PostRegisterHandler) *PostRegister {
	return &PostRegister{Context: ctx, Handler: handler}
}

/*
	PostRegister swagger:route POST /register postRegister

Register user by username and password
*/
type PostRegister struct {
	Context *middleware.Context
	Handler PostRegisterHandler
}

func (o *PostRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostRegisterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostRegisterBody post register body
//
// swagger:model PostRegisterBody
type PostRegisterBody struct {

	// login
	// Required: true
	Login *string `json:"login"`

	// password
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this post register body
func (o *PostRegisterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostRegisterBody) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"login", "body", o.Login); err != nil {
		return err
	}

	return nil
}

func (o *PostRegisterBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post register body based on context it is used
func (o *PostRegisterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRegisterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRegisterBody) UnmarshalBinary(b []byte) error {
	var res PostRegisterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRegisterOKBody post register o k body
//
// swagger:model PostRegisterOKBody
type PostRegisterOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this post register o k body
func (o *PostRegisterOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post register o k body based on context it is used
func (o *PostRegisterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRegisterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRegisterOKBody) UnmarshalBinary(b []byte) error {
	var res PostRegisterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}