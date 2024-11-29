package handlers

import (
	"auth/internal/impl"
	"auth/internal/models"
	"auth/internal/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) RegisterHandler(api operations.PostRegisterParams) middleware.Responder {
	token, err := impl.CreateUser(h.Client, api.Body)
	if err != nil {
		conflict := int64(operations.PostRegisterConflictCode)
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	result := new(operations.PostRegisterOK).WithPayload(&operations.PostRegisterOKBody{
		Token: *token,
	})
	return result
}
