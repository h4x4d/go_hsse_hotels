package handlers

import (
	"auth/internal/impl"
	"auth/internal/models"
	"auth/internal/restapi/operations"
	"context"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) ChangePasswordHandler(params operations.PostChangePasswordParams) middleware.Responder {
	_, span := h.tracer.Start(context.Background(), "change_password")
	defer span.End()
	token, err := impl.ChangePasswordUser(h.Client, params.Body)
	conflict := int64(operations.PostLoginUnauthorizedCode)
	if err != nil {
		return new(operations.PostChangePasswordUnauthorized).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	result := new(operations.PostChangePasswordOK).WithPayload(&operations.PostChangePasswordOKBody{
		Token: *token,
	})
	return result
}
