package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
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
