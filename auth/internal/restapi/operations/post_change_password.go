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

// PostChangePasswordHandlerFunc turns a function with the right signature into a post change password handler
type PostChangePasswordHandlerFunc func(PostChangePasswordParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostChangePasswordHandlerFunc) Handle(params PostChangePasswordParams) middleware.Responder {
	return fn(params)
}

// PostChangePasswordHandler interface for that can handle valid post change password params
type PostChangePasswordHandler interface {
	Handle(PostChangePasswordParams) middleware.Responder
}

// NewPostChangePassword creates a new http.Handler for the post change password operation
func NewPostChangePassword(ctx *middleware.Context, handler PostChangePasswordHandler) *PostChangePassword {
	return &PostChangePassword{Context: ctx, Handler: handler}
}

/*
	PostChangePassword swagger:route POST /change-password postChangePassword

Change password
*/
type PostChangePassword struct {
	Context *middleware.Context
	Handler PostChangePasswordHandler
}

func (o *PostChangePassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostChangePasswordParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostChangePasswordBody post change password body
//
// swagger:model PostChangePasswordBody
type PostChangePasswordBody struct {

	// login
	// Required: true
	Login *string `json:"login"`

	// new password
	// Required: true
	NewPassword *string `json:"newPassword"`

	// old password
	// Required: true
	OldPassword *string `json:"oldPassword"`
}

// Validate validates this post change password body
func (o *PostChangePasswordBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOldPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostChangePasswordBody) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"login", "body", o.Login); err != nil {
		return err
	}

	return nil
}

func (o *PostChangePasswordBody) validateNewPassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"newPassword", "body", o.NewPassword); err != nil {
		return err
	}

	return nil
}

func (o *PostChangePasswordBody) validateOldPassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"oldPassword", "body", o.OldPassword); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post change password body based on context it is used
func (o *PostChangePasswordBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostChangePasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChangePasswordBody) UnmarshalBinary(b []byte) error {
	var res PostChangePasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChangePasswordOKBody post change password o k body
//
// swagger:model PostChangePasswordOKBody
type PostChangePasswordOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this post change password o k body
func (o *PostChangePasswordOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post change password o k body based on context it is used
func (o *PostChangePasswordOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostChangePasswordOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChangePasswordOKBody) UnmarshalBinary(b []byte) error {
	var res PostChangePasswordOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
