package handlers

import (
    "auth/internal/models"
    "auth/internal/restapi/operations"
    "github.com/go-openapi/runtime/middleware"
)

func (h *Handler) ChangePasswordHandler(params operations.PostChangePasswordParams) middleware.Responder {
    token, err := h.Client.ChangePasswordUser(params.Body)
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