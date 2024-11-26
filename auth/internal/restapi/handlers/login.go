package handlers

import (
	"auth/internal/models"
	"auth/internal/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) LoginHandler(api operations.PostLoginParams) middleware.Responder {
	token, err := h.Client.LoginUser(api.Body)
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
