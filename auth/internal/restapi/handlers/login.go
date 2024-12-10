package handlers

import (
	"auth/internal/impl"
	"auth/internal/models"
	"auth/internal/restapi/operations"
	"context"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) LoginHandler(api operations.PostLoginParams) middleware.Responder {
	_, span := h.tracer.Start(context.Background(), "login")
	defer span.End()
	token, err := impl.LoginUser(h.Client, api.Body)
	conflict := int64(operations.PostLoginUnauthorizedCode)
	if err != nil {
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	result := new(operations.PostLoginOK).WithPayload(&operations.PostLoginOKBody{
		Token: *token,
	})
	return result
}
