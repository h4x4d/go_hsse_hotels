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

// PostLoginHandlerFunc turns a function with the right signature into a post login handler
type PostLoginHandlerFunc func(PostLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLoginHandlerFunc) Handle(params PostLoginParams) middleware.Responder {
	return fn(params)
}

// PostLoginHandler interface for that can handle valid post login params
type PostLoginHandler interface {
	Handle(PostLoginParams) middleware.Responder
}

// NewPostLogin creates a new http.Handler for the post login operation
func NewPostLogin(ctx *middleware.Context, handler PostLoginHandler) *PostLogin {
	return &PostLogin{Context: ctx, Handler: handler}
}

/*
	PostLogin swagger:route POST /login postLogin

Sign in user by login and password
*/
type PostLogin struct {
	Context *middleware.Context
	Handler PostLoginHandler
}

func (o *PostLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostLoginBody post login body
//
// swagger:model PostLoginBody
type PostLoginBody struct {

	// login
	// Required: true
	Login *string `json:"login"`

	// password
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this post login body
func (o *PostLoginBody) Validate(formats strfmt.Registry) error {
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

func (o *PostLoginBody) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"login", "body", o.Login); err != nil {
		return err
	}

	return nil
}

func (o *PostLoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post login body based on context it is used
func (o *PostLoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginBody) UnmarshalBinary(b []byte) error {
	var res PostLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostLoginOKBody post login o k body
//
// swagger:model PostLoginOKBody
type PostLoginOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this post login o k body
func (o *PostLoginOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post login o k body based on context it is used
func (o *PostLoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginOKBody) UnmarshalBinary(b []byte) error {
	var res PostLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
